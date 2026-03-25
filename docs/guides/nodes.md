# Nodes

> Exposes node-scoped operations and acts as entry point for node-scoped services (LXC, Tasks, ...).

## Table of Contents

- [Prerequisites](#prerequisites)
- [GetNodes](#getnodes)
- [See Also](#see-also)

## Prerequisites

Requires an initialized client. See the [Client documentation](../getting-started/authentication.md).

## GetNodes

Returns all nodes in the cluster.  
**Proxmox API:** [`GET /nodes`](https://pve.proxmox.com/pve-docs/api-viewer/#/nodes)

**Go**
```go
func (c *Client) GetNodes() (*types.NodesResponse, error)
```
```go
nodes, err := client.GetNodes()
if err != nil {
    return fmt.Errorf("nodes.GetNodes: %w", err)
}

for _, node := range nodes.Data {
    fmt.Println(node.Node)
}
```

Returns [`*types.NodesResponse`](/types/nodes.go).
> [!NOTE]
> Requests time out after 30 seconds. A timeout returns a wrapped error. Check with `errors.Is(err, context.DeadlineExceeded)`.

**TypeScript**
```typescript
// Work In Progress
```

---

> [!IMPORTANT]
> `client.Node("node-name")` initializes a node-scoped service. The node name must match exactly the name returned by `GetNodes()`.

---

## See Also

- [Client documentation](../getting-started/authentication.md)
- [Other Guides](/docs/guides/)
- [Nodes associated types](/types/nodes.go)
- [Proxmox API - Nodes endpoints](https://pve.proxmox.com/pve-docs/api-viewer/#/nodes)
- [Proxmox Wiki - Proxmox Node Management](https://pve.proxmox.com/wiki/Node_management)