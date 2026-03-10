package lxc

import (
	"context"
	"fmt"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/internal/nodes/lxc/status"
	"github.com/HavelCTF/ProxmoxSDK/types"
	"github.com/google/go-querystring/query"
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

func (s *LXCService) CloneLXC(data types.CloneLXCData) (*types.CloneLXCResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	finalData := CloneLXCFinalData{
		CloneLXCData: data,
		Node:         s.node,
		VMID:         s.vmid,
	}
	values, err := query.Values(finalData)
	if err != nil {
		return nil, fmt.Errorf("[%s] Failed to query request: %w", s.c.UUID(), err)
	}

	return http.DoRequest[types.CloneLXCResponse](ctx, s.c,
		http.RequestContent{
			Method:   "POST",
			Endpoint: fmt.Sprintf("/nodes/%s/lxc/%d/clone", s.node, s.vmid),
			Body:     &values,
		},
	)
}

func (s *LXCService) Status() *status.StatusService {
	return status.New(
		status.StatusContext{
			Client: s.c,
			Node:   s.node,
			VMID:   s.vmid,
		},
	)
}
