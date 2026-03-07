# Cluster

> Cluster service exposes cluster tools from Proxmox API (in progress tasks, nextid, ...).

## Table of Contents

- [Prerequisites](#prerequisites)
- [Go](#go)
  - [Initialization](#initialization)
  - [GetTasks](#gettasks)
  - [GetNextID](#getnextid)
- [TypeScript](#typescript)
  - [Initialization](#initialization-1)
  - [GetTasks](#gettasks-1)
  - [GetNextID](#getnextid-1)
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

### GetNextID

> Return next available VMID for LXC / VM creation.
```go
// ... Client initialization

nextID, err := cluster.GetNextId() // *types.ClusterNextIdResponse, error

// OR

nextID, err := client.Cluster().GetNextID()
```

> See [`ClusterNextIdResponse`](../types/cluster.go) for the full list of available fields.

## TypeScript

### Initialization

> Work In Progress

### GetTasks

> Work In Progress

### GetNextID

> Work In Progress

## See Also

<details>
<summary>Previous</summary>

- [Client](client.md)

</details>