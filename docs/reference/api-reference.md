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
static async create(host: string, token: string, uuid: string, options?: ProxmoxSDKOptions): Promise<ProxmoxSDK>
```

| Parameter | Type | Description |
|-----------|------|-------------|
| `baseURL` | `string` | Proxmox instance URL (e.g. `https://your-host:8006`) |
| `token` | `string` | Proxmox API token (`PVEAPIToken=user@realm!name=uuid`) |
| `uuid` | `string` | Your app UUID for logs |
| `opts` / `options` | `...ClientOption` / `ProxmoxSDKOptions` | Optional configuration |

**Available options**

| Option | Go | TypeScript |
|--------|----|------------|
| Override HTTP client | `WithHTTPClient(c *http.Client)` | / |
| Disable SSL verification | / | `insecure?: boolean` |

### Service Accessors

**Go**
```go
func (c *Client) Cluster() *cluster.ClusterService
func (c *Client) Node(node string) *nodes.NodeService
```

**TypeScript**
```typescript
client.cluster(): ClusterService
client.node(nodeName: string): NodeService
```

### Methods

**Go**
```go
func (c *Client) GetNodes() (*types.NodesResponse, error)
func (c *Client) GetVersion() (*types.VersionResponse, error)
```

**TypeScript**
```typescript
async getVersion(): Promise<VersionResponse>
async getNodes(): Promise<NodesResponse>
```

> [!WARNING]
> Promises may reject with an `Error` if the request fails. Handle rejections with `try/catch` or check `instanceof Error` on the resolved value.
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
async getTasks(): Promise<ClusterNextIdResponse>
async getNextId(): Promise<ClusterTasksResponse>
```

> [!WARNING]
> Promises may reject with an `Error` if the request fails. Handle rejections with `try/catch` or check `instanceof Error` on the resolved value.

---

## NodeService

Initialized via `client.Node(node string)` / `client.node(nodeName: string)`.

### Service Accessors

**Go**
```go
func (s *NodeService) Tasks(upid string) *tasks.TaskService
func (s *NodeService) LXC(vmid int) *lxc.LXCService
```

**TypeScript**
```typescript
tasks(upid: string): TaskService
lxc(vmid: number): LXCService
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
async getLXCs(): Promise<LXCsResponse>
async getTasks(): Promise<NodeTasksResponse>
async postLXC(data: CreateLXCData): Promise<TaskBaseResponse>
```

> [!WARNING]
> Promises may reject with an `Error` if the request fails. Handle rejections with `try/catch` or check `instanceof Error` on the resolved value.

---

## TaskService

Initialized via `client.Node(node).Tasks(upid string)` / `client.node(nodeName).tasks(upid: string)`.

**Go**
```go
func (s *TaskService) GetTaskStatus() (*types.NodeTaskStatusResponse, error)
func (s *TaskService) DeleteTask() (*types.NodeTaskDeleteResponse, error)
```

**TypeScript**
```typescript
async getTaskStatus(): Promise<NodeTaskStatusResponse>
async deleteTask(): Promise<void>
```

> [!WARNING]
> Promises may reject with an `Error` if the request fails. Handle rejections with `try/catch` or check `instanceof Error` on the resolved value.

---

## LXCService

Initialized via `client.Node(node).LXC(vmid int)` / `client.node(nodeName).lxc(vmid: number)`.

### Service Accessors

**Go**
```go
func (s *LXCService) Status() *status.StatusService
```

**TypeScript**
```typescript
status(): LXCStatusService 
```

### Methods

**Go**
```go
func (s *LXCService) CloneLXC(data types.CloneLXCData) (*types.CloneLXCResponse, error)
func (s *LXCService) DeleteLXC() (*types.DeleteLXCResponse, error)
```

**TypeScript**
```typescript
async cloneLXC(data: CloneLXCData): Promise<TaskBaseResponse>
async deleteLXC(): Promise<TaskBaseResponse>
```

> [!WARNING]
> Promises may reject with an `Error` if the request fails. Handle rejections with `try/catch` or check `instanceof Error` on the resolved value.

---

## StatusService

Initialized via `client.Node(node).LXC(vmid).Status()` / `client.node(nodeName).lxc(vmid).status()`.

**Go**
```go
func (s *StatusService) StartLXC() (*types.StartLXCResponse, error)
func (s *StatusService) StopLXC() (*types.StopLXCResponse, error)
```

**TypeScript**
```typescript
async startLXC(): Promise<TaskBaseResponse>
async stopLXC(): Promise<TaskBaseResponse>
```

> [!WARNING]
> Promises may reject with an `Error` if the request fails. Handle rejections with `try/catch` or check `instanceof Error` on the resolved value.

---

## See Also

- [All Guides](../guides/)
- [Types Reference](./types.md)
- [Errors Reference](./errors.md)