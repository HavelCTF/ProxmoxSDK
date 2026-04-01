package proxmox

import (
	"net/http"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
)

type Client struct {
	c *client.Client
}

type ClientOption func(*Client)

// NewClient creates an HTTP client for the Proxmox API with the specified
// API URL, authorization token, and UUID for logging.
func NewClient(baseURL string, token string, uuid string, opts ...ClientOption) *Client {
	client := &Client{
		c: client.NewClient(baseURL, token, uuid),
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.c.SetHTTPClient(httpClient)
	}
}

func (c *Client) UUID() string {
	return c.c.UUID()
}

func (c *Client) BaseURL() string {
	return c.c.BaseURL()
}
