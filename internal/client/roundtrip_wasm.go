//go:build js && wasm

package client

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"syscall/js"
)

var uint8Array = js.Global().Get("Uint8Array")

// fetchTransport is an http.RoundTripper that uses JavaScript's fetch() API.
// Go 1.24+ disables its built-in fetch transport when running in Node.js
// (see go.dev/issue/57613), falling back to dial which fails in WASM.
// This transport re-enables fetch-based networking for Node.js WASM usage.
type fetchTransport struct{}

func (t *fetchTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	opt := js.Global().Get("Object").New()
	opt.Set("method", req.Method)
	opt.Set("credentials", "omit")

	headers := js.Global().Get("Headers").New()
	for key, values := range req.Header {
		for _, value := range values {
			headers.Call("append", key, value)
		}
	}
	opt.Set("headers", headers)

	if req.Body != nil {
		body, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			return nil, err
		}
		if len(body) != 0 {
			buf := uint8Array.New(len(body))
			js.CopyBytesToJS(buf, body)
			opt.Set("body", buf)
		}
	}

	// Call fetch() and await the promise synchronously via a channel.
	respCh := make(chan *http.Response, 1)
	errCh := make(chan error, 1)

	var success, failure js.Func
	success = js.FuncOf(func(this js.Value, args []js.Value) any {
		defer success.Release()
		defer failure.Release()

		result := args[0]
		header := http.Header{}
		headersIt := result.Get("headers").Call("entries")
		for {
			n := headersIt.Call("next")
			if n.Get("done").Bool() {
				break
			}
			pair := n.Get("value")
			key, value := pair.Index(0).String(), pair.Index(1).String()
			header.Add(key, value)
		}

		contentLength := int64(-1)
		if cl := header.Get("Content-Length"); cl != "" {
			if parsed, err := strconv.ParseInt(cl, 10, 64); err == nil {
				contentLength = parsed
			}
		}

		statusCode := result.Get("status").Int()
		statusText := result.Get("statusText").String()

		// Read the body via a ReadableStream.
		body := result.Get("body")
		var bodyReader io.ReadCloser
		if !body.IsNull() && !body.IsUndefined() {
			bodyReader = &streamReader{stream: body.Call("getReader")}
		} else {
			bodyReader = http.NoBody
		}

		respCh <- &http.Response{
			Status:        fmt.Sprintf("%d %s", statusCode, statusText),
			StatusCode:    statusCode,
			Header:        header,
			ContentLength: contentLength,
			Body:          bodyReader,
			Request:       req,
		}
		return nil
	})

	failure = js.FuncOf(func(this js.Value, args []js.Value) any {
		defer success.Release()
		defer failure.Release()

		errCh <- fmt.Errorf("fetch failed: %s", args[0].Get("message").String())
		return nil
	})

	fetchPromise := js.Global().Call("fetch", req.URL.String(), opt)
	fetchPromise.Call("then", success, failure)

	select {
	case resp := <-respCh:
		return resp, nil
	case err := <-errCh:
		return nil, err
	case <-req.Context().Done():
		return nil, req.Context().Err()
	}
}

// streamReader reads from a JavaScript ReadableStream.
type streamReader struct {
	stream js.Value
	buf    []byte
}

func (r *streamReader) Read(p []byte) (int, error) {
	if len(r.buf) > 0 {
		n := copy(p, r.buf)
		r.buf = r.buf[n:]
		return n, nil
	}

	ch := make(chan struct {
		data []byte
		done bool
		err  error
	}, 1)

	var then, catch js.Func
	then = js.FuncOf(func(this js.Value, args []js.Value) any {
		defer then.Release()
		defer catch.Release()

		result := args[0]
		done := result.Get("done").Bool()
		if done {
			ch <- struct {
				data []byte
				done bool
				err  error
			}{nil, true, nil}
			return nil
		}
		value := result.Get("value")
		data := make([]byte, value.Get("byteLength").Int())
		js.CopyBytesToGo(data, value)
		ch <- struct {
			data []byte
			done bool
			err  error
		}{data, false, nil}
		return nil
	})

	catch = js.FuncOf(func(this js.Value, args []js.Value) any {
		defer then.Release()
		defer catch.Release()

		ch <- struct {
			data []byte
			done bool
			err  error
		}{nil, false, fmt.Errorf("stream read error: %s", args[0].String())}
		return nil
	})

	r.stream.Call("read").Call("then", then, catch)

	res := <-ch
	if res.err != nil {
		return 0, res.err
	}
	if res.done {
		return 0, io.EOF
	}

	n := copy(p, res.data)
	if n < len(res.data) {
		r.buf = res.data[n:]
	}
	return n, nil
}

func (r *streamReader) Close() error {
	ch := make(chan error, 1)

	var then, catch js.Func
	then = js.FuncOf(func(this js.Value, args []js.Value) any {
		defer then.Release()
		defer catch.Release()
		ch <- nil
		return nil
	})
	catch = js.FuncOf(func(this js.Value, args []js.Value) any {
		defer then.Release()
		defer catch.Release()
		ch <- nil // Best-effort close; don't surface errors.
		return nil
	})

	r.stream.Call("cancel").Call("then", then, catch)
	return <-ch
}

// isNodeJS reports whether the Go WASM runtime is executing inside Node.js.
func isNodeJS() bool {
	process := js.Global().Get("process")
	if process.Type() != js.TypeObject {
		return false
	}
	argv0 := process.Get("argv0")
	return argv0.Type() == js.TypeString && strings.HasPrefix(argv0.String(), "node")
}
