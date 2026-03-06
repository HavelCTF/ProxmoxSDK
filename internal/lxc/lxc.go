package lxc

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/internal/lxc/container"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

type LXCService struct {
	c    *client.Client
	node string
}

func (s *LXCService) Node() string {
	return s.node
}

func New(c *client.Client, node string) *LXCService {
	return &LXCService{
		c:    c,
		node: node,
	}
}

func (s *LXCService) GetLXCs() (*types.LXCsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.LXCsResponse](ctx, s.c,
		http.RequestContent{
			Method:   "GET",
			Endpoint: fmt.Sprintf("/nodes/%s/lxc", s.node),
		},
	)
}

func (s *LXCService) PostLXC(data types.CreateLXCData) (*types.CreateLXCResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	values, err := query.Values(data)
	if err != nil {
		return nil, fmt.Errorf("[%s] Failed to query request: %w", s.c.UUID(), err)
	}

	return http.DoRequest[types.CreateLXCResponse](ctx, s.c,
		http.RequestContent{
			Method:   "POST",
			Endpoint: fmt.Sprintf("/nodes/%s/lxc", s.node),
			Body:     &values,
		},
	)
}

func (s *LXCService) Container(vmid int) *container.ContainerService {
	return container.New(
		container.ContainerContext{
			Client: s.c,
			Node:   s.node,
			VMID:   vmid,
		},
	)
}
