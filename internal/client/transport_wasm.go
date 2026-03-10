//go:build js && wasm

package client

import "net/http"

// newPlatformHTTPClient returns an HTTP client that uses the default transport.
// In GOOS=js/GOARCH=wasm, http.DefaultTransport is a fetch-based round-tripper
// that bridges Go's net/http to the JavaScript fetch API automatically.
func newPlatformHTTPClient() *http.Client {
	return &http.Client{}
}
