package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
)

// DoMultipartRequest executes a multipart/form-data POST against the Proxmox API
// and decodes the JSON response into type R. It mirrors DoRequest's envelope but
// streams the body through a pipe — retries are disabled because the body cannot
// be replayed.
func DoMultipartRequest[R any](
	ctx context.Context,
	c *client.Client,
	endpoint string,
	fields map[string]string,
	fileField string,
	filename string,
	body io.Reader,
) (*R, error) {
	req, err := c.NewMultipartRequest(ctx, "POST", endpoint, fields, fileField, filename, body)
	if err != nil {
		return nil, err
	}

	// Bypass retryablehttp; streamed multipart bodies cannot be replayed.
	resp, err := c.DoNoRetry(req.Request)
	if err != nil {
		return nil, fmt.Errorf("[%s] (%s) Multipart request failed: %w", c.UUID(), endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("[%s] (%s) Multipart request failed with status %d: %s",
			c.UUID(), endpoint, resp.StatusCode, string(respBody))
	}

	var response R
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("[%s] failed to decode multipart response: %w", c.UUID(), err)
	}
	return &response, nil
}
