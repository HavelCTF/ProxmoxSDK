package client

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	retryhttp "github.com/hashicorp/go-retryablehttp"
)

// NewMultipartRequest builds an HTTP request whose body is a streaming
// multipart/form-data payload. The caller supplies textual fields and an
// optional file part (fileField + filename + body).
//
// Pass an empty fileField (or a nil body) to omit the file part. The auth
// header is attached the same way as Client.NewRequest.
//
// Note on retries: the request body is wrapped in a streaming PipeReader
// and cannot be replayed. Callers must run uploads through a retryablehttp
// client whose retry budget is set to 0 for this request, or convert the
// reader to a bytes buffer if retries are required.
func (c *Client) NewMultipartRequest(
	ctx context.Context,
	method string,
	endpoint string,
	fields map[string]string,
	fileField string,
	filename string,
	body io.Reader,
) (*retryhttp.Request, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL(), endpoint)

	pr, pw := io.Pipe()
	mw := multipart.NewWriter(pw)

	go func() {
		// Closing the pipe writer with a non-nil error propagates to the reader.
		var writeErr error
		defer func() {
			_ = pw.CloseWithError(writeErr)
		}()

		for k, v := range fields {
			if writeErr = mw.WriteField(k, v); writeErr != nil {
				return
			}
		}

		if fileField != "" && body != nil {
			fw, err := mw.CreateFormFile(fileField, filename)
			if err != nil {
				writeErr = err
				return
			}
			if _, err := io.Copy(fw, body); err != nil {
				writeErr = err
				return
			}
		}

		if err := mw.Close(); err != nil {
			writeErr = err
			return
		}
	}()

	// retryablehttp requires the body to implement io.ReadSeeker or to be
	// declared explicitly via FromRequest; passing the io.PipeReader as-is
	// disables retries automatically (the upload helper sets RetryMax=0).
	req, err := retryhttp.NewRequest(method, url, pr)
	if err != nil {
		_ = pw.Close()
		return nil, fmt.Errorf("[%s] failed to create multipart request: %w", c.uuid, err)
	}

	req.Header.Set("Authorization", c.apiToken)
	req.Header.Set("Content-Type", mw.FormDataContentType())
	req.Header.Set("Accept", "application/json")
	req.Request = req.Request.WithContext(ctx)
	return req, nil
}

// DoNoRetry executes the request through the underlying *http.Client only,
// bypassing the retryablehttp retry policy. Used by streaming uploads where
// the body cannot be replayed.
func (c *Client) DoNoRetry(req *http.Request) (*http.Response, error) {
	return c.httpClient.HTTPClient.Do(req)
}
