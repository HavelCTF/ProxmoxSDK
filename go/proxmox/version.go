package proxmox

import (
	"context"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/go/proxmox/types"
)

// Version retrieves the Proxmox version information with retries
func (c *Client) Version() (*types.VersionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return get[types.VersionResponse](ctx, c, "/version")
}
