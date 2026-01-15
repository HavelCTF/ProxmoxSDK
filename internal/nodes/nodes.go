package nodes

import (
	"context"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

// Nodes retrieves the Proxmox nodes informations with retries
func Get(c *client.Client) (*types.NodesResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.Get[types.NodesResponse](ctx, c, "/nodes")
}
