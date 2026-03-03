// Package tasks provides functions for the Proxmox API nodes/tasks/{upid} endpoint.
package tasks

import (
	"context"
	"fmt"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

type TaskService struct {
	c    *client.Client
	node string
	upid string
}

type TaskContext struct {
	C    *client.Client
	Node string
	UPID string
}

func New(ctx TaskContext) *TaskService {
	return &TaskService{c: ctx.C, node: ctx.Node, upid: ctx.UPID}
}

func (s *TaskService) Node() string {
	return s.node
}

func (s *TaskService) UPID() string {
	return s.upid
}

func (s *TaskService) GetTaskStatus() (*types.NodeTaskStatusResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.NodeTaskStatusResponse](ctx, s.c,
		http.RequestContent{
			Method:   "GET",
			Endpoint: fmt.Sprintf("/nodes/%s/tasks/%s/status", s.node, s.upid),
		},
	)
}

func (s *TaskService) DeleteTask() (*types.NodeTaskDeleteResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return http.DoRequest[types.NodeTaskDeleteResponse](ctx, s.c,
		http.RequestContent{
			Method:   "DELETE",
			Endpoint: fmt.Sprintf("/nodes/%s/tasks/%s", s.node, s.upid),
		},
	)
}
