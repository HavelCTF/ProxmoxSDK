package version

import (
	"context"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/go/proxmox/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/go/proxmox/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/go/proxmox/types"
)

// Version retrieves the Proxmox version information with retries
func Get(c *client.Client) (*types.VersionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.Get[types.VersionResponse](ctx, c, "/version")
}
