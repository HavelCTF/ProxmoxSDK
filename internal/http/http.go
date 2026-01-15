package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
)

func Get[T any](ctx context.Context, c *client.Client, route string) (*T, error) {
	req, err := c.NewRequest(ctx, "GET", route, nil)
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

	var getResp T
	if err := json.NewDecoder(resp.Body).Decode(&getResp); err != nil {
		err = fmt.Errorf("[%s] failed to decode version response: %w", c.GetUUID(), err)
		return nil, err
	}
	return &getResp, nil
}

// TODO Modify POST to replace usage of json by x-www-form-urlencoded
func Post[T any, R any](ctx context.Context, c *client.Client, route string, body T) (*R, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("[%s] failed to marshall body: %w", c.GetUUID(), err)
	}
	req, err := c.NewRequest(ctx, "POST", route, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		err = fmt.Errorf("[%s] (%s) Request failed: %w", c.GetUUID(), route, err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf("[%s] (%s) Request failed with status %d: %s",
			c.GetUUID(), route, resp.StatusCode, string(body))
		return nil, err
	}
	defer resp.Body.Close()

	var postResp R
	if err := json.NewDecoder(resp.Body).Decode(&postResp); err != nil {
		err = fmt.Errorf("[%s] failed to decode version response: %w", c.GetUUID(), err)
		return nil, err
	}
	return &postResp, nil
}
