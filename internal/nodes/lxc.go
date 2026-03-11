package nodes

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-querystring/query"

	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/internal/nodes/lxc"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

func (s *NodeService) GetLXCs() (*types.LXCsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.LXCsResponse](ctx, s.c,
		http.RequestContent{
			Method:   "GET",
			Endpoint: fmt.Sprintf("/nodes/%s/lxc", s.node),
		},
	)
}

func (s *NodeService) PostLXC(data types.CreateLXCData) (*types.CreateLXCResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	finalData := CreateLXCFinalData{
		Node:          s.node,
		CreateLXCData: data,
	}
	values, err := query.Values(finalData)
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

func (s *NodeService) LXC(vmid int) *lxc.LXCService {
	return lxc.New(
		lxc.LXCContext{
			Client: s.c,
			Node:   s.node,
			VMID:   vmid,
		},
	)
}
