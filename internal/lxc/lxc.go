package lxc

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	lxccontext "github.com/HavelCTF/ProxmoxSDK/internal/context/lxc"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/internal/lxc/instance"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

type LXCService struct {
	ctx lxccontext.LXCContext
}

type LXCEncoder struct {
	Data types.LXC
}

func New(c *client.Client, node string) *LXCService {
	return &LXCService{
		ctx: lxccontext.LXCContext{
			Client: c,
			Node:   node,
			VMID:   nil,
		},
	}
}

func (s *LXCService) Get() (*types.LXCInfoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.Get[types.LXCInfoResponse](ctx, s.ctx.Client, fmt.Sprintf("/nodes/%s/lxc", s.ctx.Node))
}

func (s *LXCService) Post(data types.LXC) (*types.LXCResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	payload := LXCEncoder{Data: data}

	return http.Post[types.LXCResponse](ctx, s.ctx.Client, fmt.Sprintf("/nodes/%s/lxc", s.ctx.Node), &payload)
}

func (s *LXCService) Select(vmid int) *instance.LXCInstance{
	return instance.New(s.ctx, vmid)
}

func (e *LXCEncoder) Encode() (url.Values, error) {
	data := url.Values{}
	if e.Data.Node == "" || e.Data.OSTemplate == "" || e.Data.VMID < 100 {
		return nil, fmt.Errorf("Required parameter missing:\nnode: %s\nostemplate: %s\nvmid: %d",
			e.Data.Node, e.Data.OSTemplate, e.Data.VMID)
	}
	data.Set("node", e.Data.Node)
	data.Set("ostemplate", e.Data.OSTemplate)
	data.Set("vmid", strconv.Itoa(e.Data.VMID))
	return data, nil
}
