package lxc

import (
	"context"
	"fmt"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/go/proxmox/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/go/proxmox/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/go/proxmox/types"
)

type Service struct {
	c    *client.Client
	node string
}

func New(c *client.Client, node string) *Service {
	return &Service{c: c, node: node}
}

func (s *Service) Get() (*types.LXCInfoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.Get[types.LXCInfoResponse](ctx, s.c, fmt.Sprintf("/nodes/%s/lxc", s.node))
}

// TODO Modify by encode x-www-form-urlencoded
func (s *Service) Post(data types.LXC) (*types.LXCResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	return http.Post[types.LXC, types.LXCResponse](ctx, s.c, fmt.Sprintf("/nodes/%s/lxc", s.node), data)
}
