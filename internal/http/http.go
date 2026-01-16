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

type URLRequestEncoder interface {
	Encode() (url.Values, error)
}

func Get[R any](ctx context.Context, c *client.Client, route string) (*R, error) {
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

	var getResp R
	if err := json.NewDecoder(resp.Body).Decode(&getResp); err != nil {
		err = fmt.Errorf("[%s] failed to decode version response: %w", c.GetUUID(), err)
		return nil, err
	}
	return &getResp, nil
}

func Post[R any](ctx context.Context, c *client.Client, route string, payload URLRequestEncoder) (*R, error) {
	var data = url.Values{}

	if payload != nil {
		var err error
		data, err = payload.Encode()
		if err != nil {
			err = fmt.Errorf("[%s] (%s) Failed to encode request payload: %w", c.GetUUID(), route, err)
			return nil, err
		}
	}

	req, err := c.NewRequest(ctx, "POST", route, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
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
