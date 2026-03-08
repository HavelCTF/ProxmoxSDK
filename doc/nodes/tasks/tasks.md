# Tasks

> Entry point for all task-scoped services from selected node (status, history, ...).

## Table of Contents

- [Prerequisites](#prerequisites)
- [Go](#go)
  - [Initialization](#initialization)
  - [GetTasks](#gettasks)
  - [GetTaskStatus](#gettaskstatus)
  - [DeleteTask](#deletetask)
- [TypeScript](#typescript)
- [See Also](#see-also)

## Prerequisites

Requires an initialized node service. See the [Node documentation](../nodes.md).

## Go

### Initialization

Initialize a task service from an existing node service.
```go
taskService := nodeService.Tasks("task-upid")
```

### GetTasks

Returns task history from a node.

```go
tasks, err := nodeService.GetTasks() // *types.NodeTasksResponse, error
```

See the [NodeTasksResponse](../../types/nodes.go) type for available fields.

### GetTaskStatus

Returns detailed information about a specific task.

```go
task, err := taskService.GetTaskStatus() // *types.NodeTaskStatusResponse, error

// OR

task, err := nodeService.Tasks("task-upid").GetTaskStatus() // *types.NodeTaskStatusResponse, error
```

See the [NodeTaskStatusResponse](../../types/nodes.go) type for available fields.

### DeleteTask

Stops a task.

```go
_, err := taskService.DeleteTask() // *types.NodeTaskDeleteResponse, error

// OR

_, err := nodeService.Tasks("task-upid").DeleteTask() // *types.NodeTaskDeleteResponse, error
```

In case of success, the response body is nil with HTTP status 200, no error is returned.

## TypeScript

> Work In Progress

## See Also

<details>
<summary>Previous</summary>

- [Nodes](../nodes.md)

</details>