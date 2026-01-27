// Package client provides tools to create a client capable of communicating
// with the Proxmox API using a retryable HTTP configuration.
package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"

	retryhttp "github.com/hashicorp/go-retryablehttp"
)

type Client struct {
	baseURL    string
	apiToken   string
	uuid       string
	httpClient *retryhttp.Client
}

// NewClient creates an HTTP client for Proxmox using baseUrl and apiToken.
// The client enables retryable requests and logging with uuid.
func NewClient(baseURL string, apiToken string, uuid string) *Client {
	httpClient := retryhttp.NewClient()
	httpClient.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 60 * time.Second,
	}

	httpClient.RetryMax = 3
	httpClient.RetryWaitMin = 5 * time.Second
	httpClient.RetryWaitMax = 5 * time.Second

	// Prevent retries on client errors (e.g., 401 Unauthorized). Network errors and 5xx still trigger retries.
	httpClient.ErrorHandler = func(resp *http.Response, err error, numTries int) (*http.Response, error) {
		if resp != nil {
			return resp, nil
		}
		return nil, err
	}

	// Disable retries for POST requests to avoid duplicate operations
	httpClient.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if resp != nil {
			rreq, _ := retryhttp.FromRequest(resp.Request)
			if rreq != nil && rreq.Method == http.MethodPost {
				return false, nil
			}
		}
		return retryhttp.DefaultRetryPolicy(ctx, resp, err)
	}

	baseURL = strings.TrimSuffix(baseURL, "/")

	return &Client{
		baseURL:    baseURL,
		apiToken:   apiToken,
		uuid:       uuid,
		httpClient: httpClient,
	}
}

func (c *Client) GetBaseURL() string {
	return fmt.Sprintf("%s/api2/json", c.baseURL)
}

func (c *Client) GetUUID() string {
	return c.uuid
}
