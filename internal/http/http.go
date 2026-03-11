// Package http provides a typed request executor for the Proxmox API.
// It handles request encoding (form-encoded POST), execution,
// and response decoding using a shared client.
package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
)

type RequestContent struct {
	Method   string
	Endpoint string
	Body     *url.Values
	// TODO Manage OptionalParameters (DELETE LXC)
}

func encodeBody(body *url.Values) io.Reader {
	if body == nil {
		return nil
	}
	return strings.NewReader(body.Encode())
}

// DoRequest executes a Proxmox API request and decodes the response into type R.
// POST bodies are URL-encoded.
func DoRequest[R any](ctx context.Context, c *client.Client, content RequestContent) (*R, error) {
	encodedBody := encodeBody(content.Body)

	req, err := c.NewRequest(ctx, content.Method, content.Endpoint, encodedBody)
	if err != nil {
		return nil, err
	}
	if content.Method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept", "application/json")
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[%s] (%s) Request failed: %w", c.UUID(), content.Endpoint, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("[%s] (%s) Request failed with status %d: %s",
			c.UUID(), content.Endpoint, resp.StatusCode, string(body))
	}

	var response R
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("[%s] failed to decode version response: %w", c.UUID(), err)
	}
	return &response, nil
}
