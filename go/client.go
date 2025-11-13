package proxmox

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Client represents a Proxmox API client
type Client struct {
	baseURL    string
	apiToken   string
	uuid       string
	httpClient *http.Client
}

// NewClient creates a new Proxmox client instance
func NewClient(baseURL, apiToken, uuid string) *Client {
	// Create HTTP client with insecure TLS (matching TypeScript behavior)
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 30 * time.Second,
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
