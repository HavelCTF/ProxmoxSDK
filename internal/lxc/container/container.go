package container

import (
	"context"
	"fmt"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

type ContainerService struct {
	c    *client.Client
	node string
	vmid int
}

type ContainerContext struct {
	Client *client.Client
	Node   string
	VMID   int
}

func (s *ContainerService) Node() string {
	return s.node
}

func (s *ContainerService) VMID() int {
	return s.vmid
}

func New(ctx ContainerContext) *ContainerService {
	return &ContainerService{
		c:    ctx.Client,
		node: ctx.Node,
		vmid: ctx.VMID,
	}
}

func (s *ContainerService) Delete() (*types.ContainerDeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	return http.DoRequest[types.ContainerDeleteResponse](ctx, s.c,
		http.RequestContent{
			Method:   "DELETE",
			Endpoint: fmt.Sprintf("/nodes/%s/lxc/%d", s.node, s.vmid),
		},
	)
}
