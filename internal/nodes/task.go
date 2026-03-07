package nodes

import (
	"context"
	"fmt"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/internal/nodes/tasks"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

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

func (s *NodeService) Tasks(upid string) *tasks.TaskService {
	return tasks.New(
		tasks.TaskContext{
			C:    s.c,
			Node: s.node,
			UPID: upid,
		},
	)
}
