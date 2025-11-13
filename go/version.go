package proxmox

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// VersionResponse represents the Proxmox version information
type VersionResponse struct {
	Version string `json:"version"`
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
}

// Version retrieves the Proxmox version information with retries
func (c *Client) Version() (*VersionResponse, error) {
	const maxRetries = 3
	const timeout = 10 * time.Second

	var lastErr error

	for attempt := 0; attempt < maxRetries; attempt++ {
		// Create request with timeout
		url := fmt.Sprintf("%s/version", c.GetBaseURL())
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			lastErr = fmt.Errorf("[%s] failed to create request: %w", c.uuid, err)
			continue
		}

		req.Header.Set("Authorization", c.apiToken)

		// Apply timeout to this specific request
		ctx := req.Context()
		ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		req = req.WithContext(ctxWithTimeout)

		// Execute request
		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("[%s] version request failed (attempt %d/%d): %w", c.uuid, attempt+1, maxRetries, err)
			if attempt < maxRetries-1 {
				time.Sleep(time.Duration(attempt+1) * time.Second)
			}
			continue
		}
		defer resp.Body.Close()

		// Check status code
		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			lastErr = fmt.Errorf("[%s] version request failed with status %d: %s", c.uuid, resp.StatusCode, string(body))
			if attempt < maxRetries-1 {
				time.Sleep(time.Duration(attempt+1) * time.Second)
			}
			continue
		}

		// Parse response
		var versionResp VersionResponse
		if err := json.NewDecoder(resp.Body).Decode(&versionResp); err != nil {
			lastErr = fmt.Errorf("[%s] failed to decode version response: %w", c.uuid, err)
			if attempt < maxRetries-1 {
				time.Sleep(time.Duration(attempt+1) * time.Second)
			}
			continue
		}

		return &versionResp, nil
	}

	return nil, lastErr
}
