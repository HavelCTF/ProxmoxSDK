package proxmox

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/go/proxmox/types"
)

// Version retrieves the Proxmox version information with retries
func (c *Client) Version() (*types.VersionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	req, err := initHTTPRequest(ctx, c, "GET", "/version")
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		err = fmt.Errorf("[%s] version request failed: %w", c.uuid, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf("[%s] version request failed with status %d: %s",
			c.uuid, resp.StatusCode, string(body))
		return nil, err
	}

	var versionResp types.VersionResponse
	if err := json.NewDecoder(resp.Body).Decode(&versionResp); err != nil {
		err = fmt.Errorf("[%s] failed to decode version response: %w", c.uuid, err)
		return nil, err
	}
	return &versionResp, nil
}
