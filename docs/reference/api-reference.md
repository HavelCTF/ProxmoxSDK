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
```go
func NewClient(baseURL string, token string, uuid string, opts ...ClientOption) *Client
```

| Parameter | Type | Description |
|-----------|------|-------------|
| `baseURL` | `string` | Proxmox instance URL (e.g. `https://your-host:8006`) |
| `token` | `string` | Proxmox API token (`PVEAPIToken=user@realm!name=uuid`) |
| `uuid` | `string` | Your app UUID for logs |
| `opts` | `...ClientOption` | Optional configuration |

**Available options**

| Option | Description |
|--------|-------------|
| `WithHTTPClient(c *http.Client)` | Override the HTTP client used by the retry client |

### Service Accessors
```go
func (c *Client) Cluster() *cluster.ClusterService
func (c *Client) Node(node string) *nodes.NodeService
```

### Methods
```go
func (c *Client) GetNodes() (*types.NodesResponse, error)
func (c *Client) GetVersion() (*types.VersionResponse, error)
```

---

## ClusterService

Initialized via `client.Cluster()`.
```go
func (s *ClusterService) GetNextId() (*types.ClusterNextIdResponse, error)
func (s *ClusterService) GetTasks() (*types.ClusterTasksResponse, error)
```

---

## NodeService

Initialized via `client.Node(node string)`.

### Service Accessors
```go
func (s *NodeService) Tasks(upid string) *tasks.TaskService
func (s *NodeService) LXC(vmid int) *lxc.LXCService
```

### Methods
```go
func (s *NodeService) GetLXCs() (*types.LXCsResponse, error)
func (s *NodeService) GetTasks() (*types.NodeTasksResponse, error)
func (s *NodeService) PostLXC(data types.CreateLXCData) (*types.CreateLXCResponse, error)
```

---

## TaskService

Initialized via `client.Node(node).Tasks(upid string)`.
```go
func (s *TaskService) GetTaskStatus() (*types.NodeTaskStatusResponse, error)
func (s *TaskService) DeleteTask() (*types.NodeTaskDeleteResponse, error)
```

---

## LXCService

Initialized via `client.Node(node).LXC(vmid int)`.

### Service Accessors
```go
func (s *LXCService) Status() *status.StatusService
```

### Methods
```go
func (s *LXCService) CloneLXC(data types.CloneLXCData) (*types.CloneLXCResponse, error)
func (s *LXCService) DeleteLXC() (*types.DeleteLXCResponse, error)
```

---

## StatusService

Initialized via `client.Node(node).LXC(vmid).Status()`.
```go
func (s *StatusService) StartLXC() (*types.StartLXCResponse, error)
func (s *StatusService) StopLXC() (*types.StopLXCResponse, error)
```

---

## See Also

- [All Guides](../guides/)
- [Types Reference](./types.md)
- [Errors Reference](./errors.md)