package status

import (
	"context"
	"fmt"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

type StatusService struct {
	c    *client.Client
	node string
	vmid int
}

type StatusContext struct {
	Client *client.Client
	Node   string
	VMID   int
}

func New(ctx StatusContext) *StatusService {
	return &StatusService{
		c:    ctx.Client,
		node: ctx.Node,
		vmid: ctx.VMID,
	}
}

func (s *StatusService) StartLXC() (*types.StartLXCResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.StartLXCResponse](ctx, s.c,
		http.RequestContent{
			Method:   "POST",
			Endpoint: fmt.Sprintf("/nodes/%s/lxc/%d/status/start", s.node, s.vmid),
		},
	)
}

func (s *StatusService) StopLXC() (*types.StopLXCResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.StopLXCResponse](ctx, s.c,
		http.RequestContent{
			Method:   "POST",
			Endpoint: fmt.Sprintf("/nodes/%s/lxc/%d/status/stop", s.node, s.vmid),
		},
	)
}
