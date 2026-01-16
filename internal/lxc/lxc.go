package lxc

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

type Service struct {
	c    *client.Client
	node string
}

type LXCEncoder struct {
	Data types.LXC
}

func New(c *client.Client, node string) *Service {
	return &Service{c: c, node: node}
}

func (s *Service) Get() (*types.LXCInfoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.Get[types.LXCInfoResponse](ctx, s.c, fmt.Sprintf("/nodes/%s/lxc", s.node))
}

func (s *Service) Post(data types.LXC) (*types.LXCResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	payload := LXCEncoder{Data: data}

	return http.Post[types.LXCResponse](ctx, s.c, fmt.Sprintf("/nodes/%s/lxc", s.node), &payload)
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
