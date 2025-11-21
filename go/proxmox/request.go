package proxmox

import (
	"context"
	"fmt"

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
