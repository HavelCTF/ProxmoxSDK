package lxc

import (
	"github.com/HavelCTF/ProxmoxSDK/internal/client"
)

type LXCContext struct {
	Client *client.Client
	Node   string
	VMID   *int
}
