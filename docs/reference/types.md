# Types

> All types are defined for the ProxmoxSDK

## Table of Contents

- [Shared](#shared)
- [Cluster](#cluster)
- [Nodes](#nodes)
- [LXC](#lxc)
- [Version](#version)
- [See Also](#see-also)

---

## Shared

**Go**

| Type | Description |
|------|-------------|
| `TaskBaseResponse` | Returned by all asynchronous operations. Contains the `UPID` to track task progress |

**TypeScript**

| Type | Description |
|------|-------------|
| `TaskBaseResponse` | Returned by all asynchronous operations. Contains the `UPID` to track task progress. It is validated via `TaskBaseResponseSchema`|

---

## Cluster

**Go**

| Type | Description |
|------|-------------|
| `ClusterNextIdResponse` | Returned by `Cluster.GetNextId()` |
| `ClusterTasksResponse` | Returned by `Cluster.GetTasks()` |

**TypeScript**

| Type | Description |
|------|-------------|
| `ClusterNextIdResponse` | Returned by `cluster.getNextId()` and validated via `ClusterNextIdResponseSchema` |
| `ClusterTasksResponse` | Returned by `cluster.getTasks()` and validated via `ClusterTasksResponseSchema` |

---

## Nodes

**Go**

| Type | Description |
|------|-------------|
| `NodesResponse` | Returned by `Client.GetNodes()` |
| `NodeTasksResponse` | Returned by `NodeService.GetTasks()` |
| `NodeTaskStatusResponse` | Returned by `TaskService.GetTaskStatus()` |
| `NodeTaskDeleteResponse` | Returned by `TaskService.DeleteTask()` |

**TypeScript**

| Type | Description |
|------|-------------|
| `NodesResponse` | Returned by `client.getNodes()` and validated via `NodesResponseSchema` |
| `NodeTasksResponse` | Returned by `node.getTasks()` and validated via `NodeTasksResponseSchema` |
| `NodeTaskStatusResponse` | Returned by `tasks.getTaskStatus()` and validated via `NodeTaskStatusResponseSchema` |
| `void` | Returned by `tasks.deleteTask()` |

---

## LXC

**Go**

| Type | Description |
|------|-------------|
| `LXCsResponse` | Returned by `NodeService.GetLXCs()` |
| `CreateLXCData` | Request body for `NodeService.PostLXC()` |
| `CreateLXCResponse` | Returned by `NodeService.PostLXC()`, alias of `TaskBaseResponse` |
| `CloneLXCData` | Request body for `LXCService.CloneLXC()` |
| `CloneLXCResponse` | Returned by `LXCService.CloneLXC()`, alias of `TaskBaseResponse` |
| `DeleteLXCResponse` | Returned by `LXCService.DeleteLXC()`, alias of `TaskBaseResponse` |
| `StartLXCResponse` | Returned by `StatusService.StartLXC()`, alias of `TaskBaseResponse` |
| `StopLXCResponse` | Returned by `StatusService.StopLXC()`, alias of `TaskBaseResponse` |

**TypeScript**

| Type | Description |
|------|-------------|
| `LXCsResponse` | Returned by `node.getLXCs()` and validated via `LXCsResponseSchema` |
| `CreateLXCData` | Request body for `node.postLXC()` |
| `CloneLXCData` | Request body for `lxc.cloneLXC()` |
| `TaskBaseResponse` | Returned by `node.postLXC()`, `lxc.cloneLXC()`, `lxc.deleteLXC()`, `lxc.status().startLXC()`, `lxc.status().stopLXC()`. See [Shared](#shared)|
---

## Version

**Go**

| Type | Description |
|------|-------------|
| `VersionResponse` | Returned by `Client.GetVersion()` |

**TypeScript**

| Type | Description |
|------|-------------|
| `VersionResponse` | Returned by `client.getVersion()` and validated via `VersionResponseSchema` |

---

## See Also

- [API Reference](./api-reference.md)
- [Errors Reference](./errors.md)
- [Go type definitions](/types/)
- [TypeScript type definitions](/pkg/typescript/src/types/)
- [Proxmox API viewer](https://pve.proxmox.com/pve-docs/api-viewer/)