//go:build js && wasm

package client

import (
	"net/http"
	"time"
)

// newPlatformHTTPClient returns an HTTP client for the WASM target.
// Go 1.24+ disables its built-in fetch transport when detecting Node.js,
// so we use our own fetchTransport that calls JavaScript's fetch() directly.
// In browsers, Go's DefaultTransport already uses fetch, but we use our
// transport unconditionally for consistency.
func newPlatformHTTPClient() *http.Client {
	return &http.Client{
		Transport: &fetchTransport{},
		Timeout:   60 * time.Second,
	}
}
