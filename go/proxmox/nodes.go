package proxmox

import (
	"context"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/go/proxmox/types"
)

// Nodes retrieves the Proxmox nodes informations with retries
func (c *Client) Nodes() (*types.NodesResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return get[types.NodesResponse](ctx, c, "/nodes")
}
