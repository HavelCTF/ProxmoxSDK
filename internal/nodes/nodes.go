// Package version provides functions for the Proxmox API nodes endpoint.
package nodes

import (
	"context"
	"fmt"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

type NodeService struct {
	c    *client.Client
	node string
}

func New(c *client.Client, node string) *NodeService {
	return &NodeService{c: c, node: node}
}

func (s *NodeService) Node() string {
	return s.node
}

// GetNode retrieves response of /nodes/{node} endpoint.
func (s *NodeService) GetNode() (*types.NodeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.NodeResponse](ctx, s.c,
		http.RequestContent{
			Method:   "GET",
			Endpoint: fmt.Sprintf("/nodes/%s", s.node),
		},
	)
}

// GetNodes retrieves response of /nodes endpoint.
func GetNodes(c *client.Client) (*types.NodesResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.NodesResponse](ctx, c,
		http.RequestContent{
			Method:   "GET",
			Endpoint: "/nodes",
		},
	)
}
