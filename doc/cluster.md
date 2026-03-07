# Cluster

> Cluster service exposes cluster tools from Proxmox API (in progress tasks, nextid, ...).

## Table of Contents

- [Prerequisites](#prerequisites)
- [Go](#go)
  - [Initialization](#initialization)
  - [GetTasks](#gettasks)
  - [GetNextId](#getnextid)
- [TypeScript](#typescript)
  - [Initialization](#initialization-1)
  - [GetTasks](#gettasks-1)
  - [GetNextId](#getnextid-1)
- [See Also](#see-also)

## Prerequisites

- Initialized client. See [Client](client.md)

## Go

### Initialization
```go
// ... Client initialization

cluster := client.Cluster()
```

### GetTasks

> Return all in progress tasks.
```go
// ... Client initialization

tasks, err := cluster.GetTasks() // *types.ClusterTasksResponse, error

// OR

tasks, err := client.Cluster().GetTasks()
```

> See [`ClusterTasksResponse`](../types/cluster.go) for the full list of available fields.

### GetNextId

> Return next available VMID for LXC / VM creation.
```go
// ... Client initialization

nextId, err := cluster.GetNextId() // *types.ClusterNextIdResponse, error

// OR

nextId, err := client.Cluster().GetNextId()
```

> See [`ClusterNextIdResponse`](../types/cluster.go) for the full list of available fields.

## TypeScript

### Initialization

> Work In Progress

### GetTasks

> Work In Progress

### GetNextId

> Work In Progress

## See Also

<details>
<summary>Previous</summary>

- [Client](client.md)

</details>