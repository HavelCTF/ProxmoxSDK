// Package version provides functions for the Proxmox API version endpoint.
package version

import (
	"context"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

func GetVersion(c *client.Client) (*types.VersionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.VersionResponse](ctx, c,
		http.RequestContent{
			Method:   "GET",
			Endpoint: "/version",
		},
	)
}
