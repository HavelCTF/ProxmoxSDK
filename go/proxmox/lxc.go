package proxmox

import (
	"context"
	"fmt"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/go/proxmox/types"
)

func (c *Client) GetLXC(nodeName string) (*types.GetLXCResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return get[types.GetLXCResponse](ctx, c, fmt.Sprintf("/nodes/%s/lxc", nodeName))
}
