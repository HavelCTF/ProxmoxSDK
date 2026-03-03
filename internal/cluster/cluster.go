// Package cluster provides functions for the Proxmox API /cluster endpoint.
package cluster

import (
	"context"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

type ClusterService struct {
	c *client.Client
}

func New(c *client.Client) *ClusterService {
	return &ClusterService{c: c}
}

func (s *ClusterService) GetNextId() (*types.ClusterNextIdResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.ClusterNextIdResponse](ctx, s.c,
		http.RequestContent{
			Method:   "GET",
			Endpoint: "/cluster/nextid",
		},
	)
}

func (s *ClusterService) GetTasks() (*types.ClusterTasksResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.ClusterTasksResponse](ctx, s.c,
		http.RequestContent{
			Method:   "GET",
			Endpoint: "/cluster/tasks",
		},
	)
}
