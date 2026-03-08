package lxc

import (
	"context"
	"fmt"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

type LXCService struct {
	c    *client.Client
	node string
	vmid int
}

type LXCContext struct {
	Client *client.Client
	Node   string
	VMID   int
}

func (s *LXCService) Node() string {
	return s.node
}

func (s *LXCService) VMID() int {
	return s.vmid
}

func New(ctx LXCContext) *LXCService {
	return &LXCService{
		c:    ctx.Client,
		node: ctx.Node,
		vmid: ctx.VMID,
	}
}

func (s *LXCService) DeleteLXC() (*types.DeleteLXCResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	return http.DoRequest[types.DeleteLXCResponse](ctx, s.c,
		http.RequestContent{
			Method:   "DELETE",
			Endpoint: fmt.Sprintf("/nodes/%s/lxc/%d", s.node, s.vmid),
		},
	)
}
