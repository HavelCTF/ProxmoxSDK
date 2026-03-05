package lxc

import (
	"context"
	"fmt"
	"time"

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

func (s *LXCService) Container(vmid int) *container.ContainerService {
	return container.New(
		container.ContainerContext{
			Client: s.c,
			Node:   s.node,
			VMID:   vmid,
		},
	)
}

// type LXCEncoder struct {
// 	Data types.LXC
// }

// func (s *LXCService) Post(data types.LXC) (*types.LXCResponse, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
// 	defer cancel()
// 	payload := LXCEncoder{Data: data}

// 	return http.DoRequest[types.LXCResponse](ctx, s.ctx.Client,
// 		http.RequestContent{
// 			Method:   "POST",
// 			Endpoint: fmt.Sprintf("/nodes/%s/lxc", s.ctx.Node),
// 			Body:     &payload,
// 		},
// 	)
// }

// func (e *LXCEncoder) Encode() (url.Values, error) {
// 	data := url.Values{}
// 	if e.Data.Node == "" || e.Data.OSTemplate == "" || e.Data.VMID < 100 {
// 		return nil, fmt.Errorf("Required parameter missing:\nnode: %s\nostemplate: %s\nvmid: %d",
// 			e.Data.Node, e.Data.OSTemplate, e.Data.VMID)
// 	}
// 	data.Set("node", e.Data.Node)
// 	data.Set("ostemplate", e.Data.OSTemplate)
// 	data.Set("vmid", strconv.Itoa(e.Data.VMID))
// 	return data, nil
// }
