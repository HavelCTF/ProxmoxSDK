// Package proxmox provides functions and services to communicate with
// the Proxmox API.
package proxmox

import (
	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/lxc"
	"github.com/HavelCTF/ProxmoxSDK/internal/nodes"
	"github.com/HavelCTF/ProxmoxSDK/internal/version"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

type Client struct {
	c *client.Client
}

// NewClient creates an HTTP client for the Proxmox API with the specified
// API URL, authorization token, and UUID for logging.
func NewClient(baseURL string, token string, uuid string) *Client {
	return &Client{
		c: client.NewClient(baseURL, token, uuid),
	}
}

func (c *Client) GetUUID() string {
	return c.c.GetUUID()
}

func (c *Client) GetBaseURL() string {
	return c.c.GetBaseURL()
}

// Version retrieves Proxmox /version endpoint.
func (c *Client) Version() (*types.VersionResponse, error) {
	return version.Get(c.c)
}

// Nodes retrieves Proxmox /nodes endpoint.
func (c *Client) Nodes() (*types.NodesResponse, error) {
	return nodes.Get(c.c)
}

// Node retrieves Proxmox /nodes/{node} endpoint.
func (c *Client) Node(node string) (*types.NodeResponse, error) {
	return nodes.GetNode(c.c, node)
}

func (c *Client) LXC(node string) *lxc.LXCService {
	return lxc.New(c.c, node)
}
