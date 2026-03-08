# Cluster

> Entry point for all cluster-scoped services (tasks, nextid, ...).

## Table of Contents

- [Prerequisites](#prerequisites)
- [Go](#go)
  - [Initialization](#initialization)
  - [GetTasks](#gettasks)
  - [GetNextId](#getnextid)
- [TypeScript](#typescript)
- [See Also](#see-also)

## Prerequisites

Requires an initialized client. See the [Client documentation](../client.md).

## Go

### Initialization

Initialize a cluster service from an existing client.
```go
cluster := client.Cluster()
```

### GetTasks

Return all in progress tasks.

```go
tasks, err := cluster.GetTasks() // *types.ClusterTasksResponse, error

// OR

tasks, err := client.Cluster().GetTasks() // *types.ClusterTasksResponse, error
```

See the [ClusterTasksResponse](../../types/cluster.go) type for available fields.

### GetNextId

Return next available VMID for LXC / VM creation.

```go
nextId, err := cluster.GetNextId() // *types.ClusterNextIdResponse, error

// OR

nextId, err := client.Cluster().GetNextId() // *types.ClusterNextIdResponse, error
```

See the [ClusterNextIdResponse](../../types/cluster.go) type for available fields.

## TypeScript

> Work In Progress

## See Also

<details>
<summary>Previous</summary>

- [Client](../client.md)

</details>