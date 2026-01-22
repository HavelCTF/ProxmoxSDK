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

type RequestContent struct {
	Method string
	Route  string
	Body   URLRequestEncoder
	//TODO Manage OptionalParameters (DELETE LXC)
}

func getPayload(payload URLRequestEncoder) (io.Reader, error) {
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

func DoRequest[R any](ctx context.Context, c *client.Client, content RequestContent) (*R, error) {
	payload, err := getPayload(content.Body)
	if err != nil {
		return nil, fmt.Errorf("[%s] (%s) %w", c.GetUUID(), content.Route, err)
	}

	req, err := c.NewRequest(ctx, content.Method, content.Route, payload)
	if err != nil {
		return nil, err
	}
	if content.Method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept", "application/json")
	}

	resp, err := c.Do(req)
	if err != nil {
		err = fmt.Errorf("[%s] (%s) Request failed: %w", c.GetUUID(), content.Route, err)
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf("[%s] (%s) Request failed with status %d: %s",
			c.GetUUID(), content.Route, resp.StatusCode, string(body))
		return nil, err
	}

	var response R
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		err = fmt.Errorf("[%s] failed to decode version response: %w", c.GetUUID(), err)
		return nil, err
	}
	return &response, nil

}
