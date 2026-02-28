// Package version provides functions for the Proxmox API nodes endpoint.
package nodes

import (
	"context"
	"fmt"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/internal/nodes/tasks"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

type NodeService struct {
	c    *client.Client
	node string
}

func (s *NodeService) Node() string {
	return s.node
}

func New(c *client.Client, node string) *NodeService {
	return &NodeService{c: c, node: node}
}

// GetNode retrieves response of /nodes/{node} endpoint.
func (s *NodeService) GetNode() (*types.NodeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.NodeResponse](ctx, s.c,
		http.RequestContent{
			Method:   "GET",
			Endpoint: fmt.Sprintf("/nodes/%s", s.node),
		},
	)
}

// GetNodes retrieves response of /nodes endpoint.
func GetNodes(c *client.Client) (*types.NodesResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.NodesResponse](ctx, c,
		http.RequestContent{
			Method:   "GET",
			Endpoint: "/nodes",
		},
	)
}

// GetTasks retrieves response of /nodes/{node}/tasks endpoint.
func (s *NodeService) GetTasks() (*types.NodeTasksResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.NodeTasksResponse](ctx, s.c,
		http.RequestContent{
			Method:   "GET",
			Endpoint: fmt.Sprintf("/nodes/%s/tasks", s.node),
		},
	)
}

func (s *NodeService) Tasks(UPID string) *tasks.TaskService {
	return tasks.New(
		tasks.TaskContext{
			C:    s.c,
			Node: s.node,
			UPID: UPID,
		},
	)
}
