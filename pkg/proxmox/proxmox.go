// Package proxmox provides functions and services to communicate with
// the Proxmox API.
package proxmox

import (
	"github.com/HavelCTF/ProxmoxSDK/internal/cluster"
	"github.com/HavelCTF/ProxmoxSDK/internal/nodes"
	"github.com/HavelCTF/ProxmoxSDK/internal/version"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

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
