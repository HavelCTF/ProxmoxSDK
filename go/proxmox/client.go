package proxmox

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"

	retryhttp "github.com/hashicorp/go-retryablehttp"
)

// Client represents a Proxmox API client
type Client struct {
	baseURL    string
	apiToken   string
	uuid       string
	httpClient *retryhttp.Client
}

// NewClient creates a new Proxmox client instance
func NewClient(baseURL, apiToken, uuid string) *Client {
	// Create HTTP client with insecure TLS (matching TypeScript behavior)
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

	httpClient.ErrorHandler = func(resp *http.Response, err error, numTries int) (*http.Response, error) {
		if resp != nil {
			return resp, nil
		}
		return nil, err
	}

	httpClient.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if resp != nil {
			rreq, _ := retryhttp.FromRequest(resp.Request)
			if rreq != nil && rreq.Method == http.MethodPost {
				return false, nil
			}
		}
		return retryhttp.DefaultRetryPolicy(ctx, resp, err)
	}

	// Remove trailing slash from baseURL
	baseURL = strings.TrimSuffix(baseURL, "/")

	return &Client{
		baseURL:    baseURL,
		apiToken:   apiToken,
		uuid:       uuid,
		httpClient: httpClient,
	}
}

// GetBaseURL returns the base URL for API requests
func (c *Client) GetBaseURL() string {
	return fmt.Sprintf("%s/api2/json", c.baseURL)
}

// GetUUID returns the client's UUID
func (c *Client) GetUUID() string {
	return c.uuid
}
