package proxmox

import (
	"github.com/HavelCTF/ProxmoxSDK/proxmox/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/proxmox/internal/lxc"
	"github.com/HavelCTF/ProxmoxSDK/proxmox/internal/nodes"
	"github.com/HavelCTF/ProxmoxSDK/proxmox/internal/version"
	"github.com/HavelCTF/ProxmoxSDK/proxmox/types"
)

type Client struct {
	c *client.Client
}

func NewClient(baseURL string, token string, uuid string) *Client {
	return &Client{
		c: client.New(baseURL, token, uuid),
	}
}

func (c *Client) GetUUID() string {
	return c.c.GetUUID()
}

func (c *Client) GetBaseURL() string {
	return c.c.GetBaseURL()
}

func (c *Client) Version() (*types.VersionResponse, error) {
	return version.Get(c.c)
}

func (c *Client) Nodes() (*types.NodesResponse, error) {
	return nodes.Get(c.c)
}

func (c *Client) LXC(node string) *lxc.Service {
	return lxc.New(c.c, node)
}
