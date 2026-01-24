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

// Get retrieves response of /nodes endpoint.
func Get(c *client.Client) (*types.NodesResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.NodesResponse](ctx, c,
		http.RequestContent{
			Method:   "GET",
			Endpoint: "/nodes",
		},
	)
}

// GetNode retrieves response of /nodes/{node} endpoint.
func GetNode(c *client.Client, node string) (*types.NodeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.NodeResponse](ctx, c,
		http.RequestContent{
			Method:   "GET",
			Endpoint: fmt.Sprintf("/nodes/%s", node),
		},
	)
}
