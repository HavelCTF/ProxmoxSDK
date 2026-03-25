# LXC Status

> Exposes status-scoped operations for a selected LXC container (start, stop, ...).

## Table of Contents

- [Prerequisites](#prerequisites)
- [StartLXC](#startlxc)
- [StopLXC](#stoplxc)
- [See Also](#see-also)

## Prerequisites

Requires at least an initialized client. See the [Client documentation](/docs/getting-started/authentication.md).  
Requires at most an initialized LXC service. See the [LXC documentation](./lxc.md).

## StartLXC

Starts a container. Returns a UPID to track the task progress.  
**Proxmox API:** [`POST /nodes/{node}/lxc/{vmid}/status/start`](https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/status/start)

**Go**
```go
func (s *StatusService) StartLXC() (*types.StartLXCResponse, error)
```
```go
upid, err := client.Node("pve1").LXC(100).Status().StartLXC()
if err != nil {
    return fmt.Errorf("status.StartLXC: %w", err)
}

fmt.Println(upid.Data) // e.g. UPID:pve1:...
```

Returns [`*types.StartLXCResponse`](/types/lxc.go).
> [!NOTE]
> Request time out after 30 seconds. A timeout returns a wrapped error. Check with `errors.Is(err, context.DeadlineExceeded)`.

**TypeScript**
```typescript
// Work In Progress
```

> [!IMPORTANT]
> This request is asynchronous. Use the returned UPID to track task progress. See [Tasks documentation](../nodes/tasks.md).

---

## StopLXC

Stops a running container. Returns a UPID to track the task progress.  
**Proxmox API:** [`POST /nodes/{node}/lxc/{vmid}/status/stop`](https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/status/stop)

**Go**
```go
func (s *StatusService) StopLXC() (*types.StopLXCResponse, error)
```
```go
upid, err := client.Node("pve1").LXC(100).Status().StopLXC()
if err != nil {
    return fmt.Errorf("status.StopLXC: %w", err)
}

fmt.Println(upid.Data) // e.g. UPID:pve1:...
```

Returns [`*types.StopLXCResponse`](/types/lxc.go).
> [!NOTE]
> Request time out after 30 seconds. A timeout returns a wrapped error. Check with `errors.Is(err, context.DeadlineExceeded)`.


**TypeScript**
```typescript
// Work In Progress
```

> [!IMPORTANT]
> This request is asynchronous. Use the returned UPID to track task progress. See [Tasks documentation](../nodes/tasks.md).

---

## See Also

- [Client documentation](/docs/getting-started/authentication.md)
- [LXC documentation](./lxc.md)
- [Tasks documentation](../nodes/tasks.md)
- [LXC associated types](/types/lxc.go)
- [Proxmox API - LXC Status endpoints](https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/status)
- [Proxmox Wiki - Linux Container](https://pve.proxmox.com/wiki/Linux_Container)