package proxmox

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	retryhttp "github.com/hashicorp/go-retryablehttp"
)

func initHTTPRequest(ctx context.Context, c *Client, method string, path string) (*retryhttp.Request, error) {
	url := fmt.Sprintf("%s%s", c.GetBaseURL(), path)
	req, err := retryhttp.NewRequest(method, url, nil)

	if err != nil {
		err = fmt.Errorf("[%s] failed to create request: %w", c.uuid, err)
		return nil, err
	}
	req.Header.Set("Authorization", c.apiToken)
	req.Request = req.Request.WithContext(ctx)
	return req, nil
}

func get[T any](ctx context.Context, c *Client, route string) (*T, error) {
	req, err := initHTTPRequest(ctx, c, "GET", route)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		err = fmt.Errorf("[%s] (%s) Request failed: %w", c.uuid, route, err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf("[%s] (%s) Request failed with status %d: %s",
			c.uuid, route, resp.StatusCode, string(body))
		return nil, err
	}

	var nodesResp T
	if err := json.NewDecoder(resp.Body).Decode(&nodesResp); err != nil {
		err = fmt.Errorf("[%s] failed to decode version response: %w", c.uuid, err)
		return nil, err
	}
	return &nodesResp, nil
}
