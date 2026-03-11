// Package proxmox provides functions and services to communicate with
// the Proxmox API.
package proxmox

import (
	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/cluster"
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

func (c *Client) UUID() string {
	return c.c.UUID()
}

func (c *Client) BaseURL() string {
	return c.c.BaseURL()
}

func (c *Client) Cluster() *cluster.ClusterService {
	return cluster.New(c.c)
}

func (c *Client) GetNodes() (*types.NodesResponse, error) {
	return nodes.GetNodes(c.c)
}

func (c *Client) Node(node string) *nodes.NodeService {
	return nodes.New(c.c, node)
}

func (c *Client) GetVersion() (*types.VersionResponse, error) {
	return version.GetVersion(c.c)
}
