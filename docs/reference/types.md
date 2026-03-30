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
| `TaskBaseResponse` | Returned by all asynchronous operations — contains the `UPID` to track task progress |

**TypeScript**
```typescript
// Work In Progress
```

---

## Cluster

**Go**

| Type | Description |
|------|-------------|
| `ClusterNextIdResponse` | Returned by `Cluster.GetNextId()` |
| `ClusterTasksResponse` | Returned by `Cluster.GetTasks()` |

**TypeScript**
```typescript
// Work In Progress
```

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
```typescript
// Work In Progress
```

---

## LXC

**Go**

| Type | Description |
|------|-------------|
| `LXCsResponse` | Returned by `NodeService.GetLXCs()` |
| `CreateLXCData` | Request body for `NodeService.PostLXC()` |
| `CreateLXCResponse` | Returned by `NodeService.PostLXC()` — alias of `TaskBaseResponse` |
| `CloneLXCData` | Request body for `LXCService.CloneLXC()` |
| `CloneLXCResponse` | Returned by `LXCService.CloneLXC()` — alias of `TaskBaseResponse` |
| `DeleteLXCResponse` | Returned by `LXCService.DeleteLXC()` — alias of `TaskBaseResponse` |
| `StartLXCResponse` | Returned by `StatusService.StartLXC()` — alias of `TaskBaseResponse` |
| `StopLXCResponse` | Returned by `StatusService.StopLXC()` — alias of `TaskBaseResponse` |

**TypeScript**
```typescript
// Work In Progress
```

---

## Version

**Go**

| Type | Description |
|------|-------------|
| `VersionResponse` | Returned by `Client.GetVersion()` |

**TypeScript**
```typescript
// Work In Progress
```

---

## See Also

- [API Reference](./api-reference.md)
- [Errors Reference](./errors.md)
- [Go type definitions](/types/)
- [Proxmox API viewer](https://pve.proxmox.com/pve-docs/api-viewer/)