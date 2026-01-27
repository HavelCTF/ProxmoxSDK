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

// Encoder encodes Proxmox request data into URL-encoded form values.
type Encoder interface {
	Encode() (url.Values, error)
}

type RequestContent struct {
	Method   string
	Endpoint string
	Body     Encoder
	// TODO Manage OptionalParameters (DELETE LXC)
}

// encode converts the encoder's data into URL-encoded form data.
// Returns nil if enc is nil
func encodeBody(enc Encoder) (io.Reader, error) {
	var data = url.Values{}
	if enc == nil {
		return nil, nil
	}
	data, err := enc.Encode()
	if err != nil {
		return nil, fmt.Errorf("Failed to encode request payload: %w", err)
	}
	return strings.NewReader(data.Encode()), nil
}

// DoRequest executes a Proxmox API request and decodes the response into type R.
// POST bodies are URL-encoded.
func DoRequest[R any](ctx context.Context, c *client.Client, content RequestContent) (*R, error) {
	body, err := encodeBody(content.Body)
	if err != nil {
		return nil, fmt.Errorf("[%s] (%s) %w", c.GetUUID(), content.Endpoint, err)
	}

	req, err := c.NewRequest(ctx, content.Method, content.Endpoint, body)
	if err != nil {
		return nil, err
	}
	if content.Method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept", "application/json")
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[%s] (%s) Request failed: %w", c.GetUUID(), content.Endpoint, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("[%s] (%s) Request failed with status %d: %s",
			c.GetUUID(), content.Endpoint, resp.StatusCode, string(body))
	}

	var response R
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("[%s] failed to decode version response: %w", c.GetUUID(), err)
	}
	return &response, nil
}
