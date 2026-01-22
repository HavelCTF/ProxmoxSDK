package client

import (
	"context"
	"fmt"
	"net/http"

	retryhttp "github.com/hashicorp/go-retryablehttp"
)

func (c *Client) NewRequest(ctx context.Context, method string, path string, body any) (*retryhttp.Request, error) {
	url := fmt.Sprintf("%s%s", c.GetBaseURL(), path)
	req, err := retryhttp.NewRequest(method, url, body)

	if err != nil {
		err = fmt.Errorf("[%s] failed to create request: %w", c.uuid, err)
		return nil, err
	}
	req.Header.Set("Authorization", c.apiToken)
	req.Request = req.Request.WithContext(ctx)
	return req, nil
}

func (c *Client) Do(request *retryhttp.Request) (*http.Response, error) {
	return c.httpClient.Do(request)
}
