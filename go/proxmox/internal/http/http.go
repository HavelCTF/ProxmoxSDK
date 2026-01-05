package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/HavelCTF/ProxmoxSDK/go/proxmox/internal/client"
)

func Get[T any](ctx context.Context, c *client.Client, route string) (*T, error) {
	req, err := c.NewRequest(ctx, "GET", route)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		err = fmt.Errorf("[%s] (%s) Request failed: %w", c.GetUUID(), route, err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf("[%s] (%s) Request failed with status %d: %s",
			c.GetUUID(), route, resp.StatusCode, string(body))
		return nil, err
	}

	var nodesResp T
	if err := json.NewDecoder(resp.Body).Decode(&nodesResp); err != nil {
		err = fmt.Errorf("[%s] failed to decode version response: %w", c.GetUUID(), err)
		return nil, err
	}
	return &nodesResp, nil
}
