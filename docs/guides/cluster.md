# Cluster

> Entry point for cluster-scoped operations shared across all nodes.

## Table of Contents

- [Prerequisites](#prerequisites)
- [GetTasks](#gettasks)
- [GetNextId](#getnextid)
- [See Also](#see-also)

## Prerequisites

Requires an initialized client. See the [Client documentation](../getting-started/authentication.md).

## GetTasks

Returns all cluster-wide tasks. Tasks are created by any asynchronous operation on the cluster (VM start/stop, migration, backup, snapshot, ...).

**Go**
```go
func (s *ClusterService) GetTasks() (*types.ClusterTasksResponse, error)
```
```go
tasks, err := client.Cluster().GetTasks()
if err != nil {
    return fmt.Errorf("cluster.GetTasks: %w", err)
}

for _, task := range tasks.Data {
    fmt.Printf("[%s] %s on %s — %s\n", task.StartTime, task.Type, task.Node, task.Status)
}
```

Returns [`*types.ClusterTasksResponse`](/types/cluster.go).

**TypeScript**
```typescript
// Work In Progress
```

> An empty slice is a valid response. It means no tasks have been recorded yet.

---

## GetNextId

Returns the next available VMID for VM or LXC container creation.

**Go**
```go
func (s *ClusterService) GetNextId() (*types.ClusterNextIdResponse, error)
```
```go
nextId, err := client.Cluster().GetNextId()
if err != nil {
    return fmt.Errorf("cluster.GetNextId: %w", err)
}

fmt.Println(nextId.Data) // e.g. 105
```

Returns [`*types.ClusterNextIdResponse`](/types/cluster.go).

**TypeScript**
```typescript
// Work In Progress
```

> [!WARNING]
> The returned ID is **not reserved**. Another process may claim it before your creation call. Handle conflict errors on creation and re-fetch if needed.

---

## See Also

- [Client documentation](../getting-started/authentication.md)
- [Other Guides](/docs/guides/)
- [Cluster associated types](/types/cluster.go)
- [Proxmox Cluster Manager documentation](https://pve.proxmox.com/wiki/Cluster_Manager)
- [Proxmox API - Cluster endpoints](https://pve.proxmox.com/pve-docs/api-viewer/#/cluster)