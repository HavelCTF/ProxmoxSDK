# Tasks

> Exposes task-scoped operations from a selected node (status, history, stop).

## Table of Contents

- [Prerequisites](#prerequisites)
- [GetTasks](#gettasks)
- [GetTaskStatus](#gettaskstatus)
- [DeleteTask](#deletetask)
- [See Also](#see-also)

## Prerequisites

Requires at least an initialized client. See the [Client documentation](../getting-started/authentication.md).
Requires an initialized node service. See the [Nodes documentation](./nodes.md).

## GetTasks

Returns the task history from a node.  
**Proxmox API:** [`GET /nodes/{node}/tasks`](https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/tasks)

**Go**
```go
func (s *NodeService) GetTasks() (*types.NodeTasksResponse, error)
```
```go
tasks, err := client.Node("pve1").GetTasks()
if err != nil {
    return fmt.Errorf("tasks.GetTasks: %w", err)
}

for _, task := range tasks.Data {
    fmt.Printf("[%s] %s - %s\n", task.StartTime, task.Type, task.Status)
}
```

Returns [`*types.NodeTasksResponse`](/types/nodes.go).
> [!NOTE]
> Request time out after 30 seconds. A timeout returns a wrapped error. Check with `errors.Is(err, context.DeadlineExceeded)`.

**TypeScript**
```typescript
// Work In Progress
```

> An empty slice is a valid response. It means no tasks have been recorded on this node yet.

---

## GetTaskStatus

Returns detailed information about a specific task.  
**Proxmox API:** [`GET /nodes/{node}/tasks/{upid}/status`](https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/tasks/{upid}/status)

**Go**
```go
func (s *TaskService) GetTaskStatus() (*types.NodeTaskStatusResponse, error)
```
```go
task, err := client.Node("pve1").Tasks("UPID:pve1:...").GetTaskStatus()
if err != nil {
    return fmt.Errorf("tasks.GetTaskStatus: %w", err)
}

fmt.Println(task.Data.Status) // e.g. OK
```

Returns [`*types.NodeTaskStatusResponse`](/types/nodes.go).
> [!NOTE]
> Request time out after 30 seconds. A timeout returns a wrapped error. Check with `errors.Is(err, context.DeadlineExceeded)`.

**TypeScript**
```typescript
// Work In Progress
```

> [!IMPORTANT]
> `client.Node("node-name").Tasks("upid")` initializes a task-scoped service. The UPID must match exactly the value returned by `GetTasks()` or `GetTaskStatus()`.

---

## DeleteTask

Stops a running task.  
**Proxmox API:** [`DELETE /nodes/{node}/tasks/{upid}`](https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/tasks/{upid})

**Go**
```go
func (s *TaskService) DeleteTask() (*types.NodeTaskDeleteResponse, error)
```
```go
_, err := client.Node("pve1").Tasks("UPID:pve1:...").DeleteTask()
if err != nil {
    return fmt.Errorf("tasks.DeleteTask: %w", err)
}

fmt.Println("task stopped")
```

Returns [`*types.NodeTaskDeleteResponse`](/types/nodes.go).

**TypeScript**
```typescript
// Work In Progress
```

> [!WARNING]
> On success, the response body is `nil` with HTTP status 200. Discard the return value and check only the error.

---

## See Also

- [Client documentation](../getting-started/authentication.md)
- [Other Guides](/docs/guides/)
- [Cluster Tasks Guide](./cluster.md#gettasks)
- [Tasks associated types](/types/nodes.go)
- [Proxmox API - Tasks endpoints](https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/tasks)
- [Proxmox Wiki - Tasks](https://pve.proxmox.com/wiki/Proxmox_VE_API#Tasks)