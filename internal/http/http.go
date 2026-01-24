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

// Encoder encodes DTOs for Proxmox POST requests into URL-encoded values.
type Encoder interface {
	Encode() (url.Values, error)
}

type RequestContent struct {
	Method   string
	Endpoint string
	Body     Encoder
	//TODO Manage OptionalParameters (DELETE LXC)
}

// getPayload encodes the given payload using the Encoder interface.
// It returns an io.Reader containing the URL-encoded payload,
// or nil if the payload is nil.
func getPayload(payload Encoder) (io.Reader, error) {
	var data = url.Values{}
	if payload == nil {
		return nil, nil
	}
	data, err := payload.Encode()
	if err != nil {
		return nil, fmt.Errorf("Failed to encode request payload: %w", err)
	}
	return strings.NewReader(data.Encode()), nil
}

// DoRequest executes an HTTP request (GET, POST, DELETE) to the Proxmox API.
// It encodes POST request bodies as URL-encoded and decodes the response
// into the specified type R.
func DoRequest[R any](ctx context.Context, c *client.Client, content RequestContent) (*R, error) {
	payload, err := getPayload(content.Body)
	if err != nil {
		return nil, fmt.Errorf("[%s] (%s) %w", c.GetUUID(), content.Endpoint, err)
	}

	req, err := c.NewRequest(ctx, content.Method, content.Endpoint, payload)
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
