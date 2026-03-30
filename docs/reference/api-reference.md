# API Reference

> Exhaustive list of all public methods exposed by the ProxmoxSDK.

## Table of Contents

- [Client](#client)
- [ClusterService](#clusterservice)
- [NodeService](#nodeservice)
- [TaskService](#taskservice)
- [LXCService](#lxcservice)
- [StatusService](#statusservice)
- [See Also](#see-also)

---

## Client

### Initialization

**Go**
```go
func NewClient(baseURL string, token string, uuid string, opts ...ClientOption) *Client
```

**TypeScript**
```typescript
// Work In Progress
```

| Parameter | Type | Description |
|-----------|------|-------------|
| `baseURL` | `string` | Proxmox instance URL (e.g. `https://your-host:8006`) |
| `token` | `string` | Proxmox API token (`PVEAPIToken=user@realm!name=uuid`) |
| `uuid` | `string` | Your app UUID for logs |
| `opts` | `...ClientOption` | Optional configuration |

**Available options**

| Option | Go | TypeScript |
|--------|----|------------|
| Override HTTP client | `WithHTTPClient(c *http.Client)` | Work In Progress |

### Service Accessors

**Go**
```go
func (c *Client) Cluster() *cluster.ClusterService
func (c *Client) Node(node string) *nodes.NodeService
```

**TypeScript**
```typescript
client.cluster()
client.node(node: string)
```

### Methods

**Go**
```go
func (c *Client) GetNodes() (*types.NodesResponse, error)
func (c *Client) GetVersion() (*types.VersionResponse, error)
```

**TypeScript**
```typescript
// Work In Progress
```

---

## ClusterService

Initialized via `client.Cluster()` / `client.cluster()`.

**Go**
```go
func (s *ClusterService) GetNextId() (*types.ClusterNextIdResponse, error)
func (s *ClusterService) GetTasks() (*types.ClusterTasksResponse, error)
```

**TypeScript**
```typescript
// Work In Progress
```

---

## NodeService

Initialized via `client.Node(node string)` / `client.node(node: string)`.

### Service Accessors

**Go**
```go
func (s *NodeService) Tasks(upid string) *tasks.TaskService
func (s *NodeService) LXC(vmid int) *lxc.LXCService
```

**TypeScript**
```typescript
client.node(node).tasks(upid: string)
client.node(node).lxc(vmid: number)
```

### Methods

**Go**
```go
func (s *NodeService) GetLXCs() (*types.LXCsResponse, error)
func (s *NodeService) GetTasks() (*types.NodeTasksResponse, error)
func (s *NodeService) PostLXC(data types.CreateLXCData) (*types.CreateLXCResponse, error)
```

**TypeScript**
```typescript
// Work In Progress
```

---

## TaskService

Initialized via `client.Node(node).Tasks(upid string)` / `client.node(node).tasks(upid: string)`.

**Go**
```go
func (s *TaskService) GetTaskStatus() (*types.NodeTaskStatusResponse, error)
func (s *TaskService) DeleteTask() (*types.NodeTaskDeleteResponse, error)
```

**TypeScript**
```typescript
// Work In Progress
```

---

## LXCService

Initialized via `client.Node(node).LXC(vmid int)` / `client.node(node).lxc(vmid: number)`.

### Service Accessors

**Go**
```go
func (s *LXCService) Status() *status.StatusService
```

**TypeScript**
```typescript
client.node(node).lxc(vmid).status()
```

### Methods

**Go**
```go
func (s *LXCService) CloneLXC(data types.CloneLXCData) (*types.CloneLXCResponse, error)
func (s *LXCService) DeleteLXC() (*types.DeleteLXCResponse, error)
```

**TypeScript**
```typescript
// Work In Progress
```

---

## StatusService

Initialized via `client.Node(node).LXC(vmid).Status()` / `client.node(node).lxc(vmid).status()`.

**Go**
```go
func (s *StatusService) StartLXC() (*types.StartLXCResponse, error)
func (s *StatusService) StopLXC() (*types.StopLXCResponse, error)
```

**TypeScript**
```typescript
// Work In Progress
```

---

## See Also

- [All Guides](../guides/)
- [Types Reference](./types.md)
- [Errors Reference](./errors.md)