# LXC

> Exposes LXC-scoped operations from a selected node (creation, deletion, clone, ...) and acts as entry point for LXC-scoped services (Status, ...).

## Table of Contents

- [Prerequisites](#prerequisites)
- [GetLXCs](#getlxcs)
- [PostLXC](#postlxc)
- [DeleteLXC](#deletelxc)
- [CloneLXC](#clonelxc)
- [See Also](#see-also)

## Prerequisites

Requires at least an initialized client. See the [Client documentation](/docs/getting-started/authentication.md).  
Requires at most an initialized node service. See the [Nodes documentation](../nodes/nodes.md).

## GetLXCs

Returns all LXC containers from a node.  
**Proxmox API:** [`GET /nodes/{node}/lxc`](https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc)

**Go**
```go
func (s *NodeService) GetLXCs() (*types.LXCsResponse, error)
```
```go
lxcs, err := client.Node("pve1").GetLXCs()
if err != nil {
    return fmt.Errorf("lxc.GetLXCs: %w", err)
}

for _, lxc := range lxcs.Data {
    fmt.Printf("[%d] %s - %s\n", lxc.VMID, lxc.Name, lxc.Status)
}
```

Returns [`*types.LXCsResponse`](/types/lxc.go).
> [!NOTE]
> Request time out after 30 seconds. A timeout returns a wrapped error. Check with `errors.Is(err, context.DeadlineExceeded)`.

**TypeScript**
```typescript
const lxcs = await client.node("pve1").getLXCs();
if (lxcs instanceof Error) {
    console.error("node.getLXCs:", lxcs);
    return;
}

for (const lxc of lxcs.Data) {
    console.log(`[${lxc.VMID}] ${lxc.Name} - ${lxc.Status}`);
}
```

---

## PostLXC

Creates a container from a given template.  
**Proxmox API:** [`POST /nodes/{node}/lxc`](https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc)

**Go**
```go
func (s *NodeService) PostLXC(data types.CreateLXCData) (*types.CreateLXCResponse, error)
```
```go
upid, err := client.Node("pve1").PostLXC(types.CreateLXCData{
    // Mandatory
    OSTemplate: "local:vztmpl/debian-12-standard_12.12-1_amd64.tar.zst",
    VMID:       2222, // use client.Cluster().GetNextId() to get a free VMID

    // Optional
    Features: &types.LXCFeatures{
        Nesting: true,
    },
})
if err != nil {
    return fmt.Errorf("lxc.PostLXC: %w", err)
}

fmt.Println(upid.Data) // e.g. UPID:pve1:...
```

Returns [`*types.CreateLXCResponse`](/types/lxc.go).
> [!NOTE]
> Request time out after 30 seconds. A timeout returns a wrapped error. Check with `errors.Is(err, context.DeadlineExceeded)`.

**TypeScript**
```typescript
const upid = await client.node("pve1").postLXC({
    // Mandatory
    OSTemplate: "local:vztmpl/debian-12-standard_12.12-1_amd64.tar.zst",
    VMID: 2222, // use client.cluster().getNextId() to get a free VMID

    // Optional
    Features: {
        Nesting: true,
    },
});
if (upid instanceof Error) {
    console.error("node.postLXC:", upid);
    return;
}

console.log(upid.Data); // e.g. UPID:pve1:...
```

> [!IMPORTANT]
> This request is asynchronous. Use the returned UPID to track task progress. See [Tasks documentation](../nodes/tasks.md).

---

## DeleteLXC

Deletes a container.  
**Proxmox API:** [`DELETE /nodes/{node}/lxc/{vmid}`](https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid})

**Go**
```go
func (s *LXCService) DeleteLXC() (*types.DeleteLXCResponse, error)
```
```go
upid, err := client.Node("pve1").LXC(2222).DeleteLXC()
if err != nil {
    return fmt.Errorf("lxc.DeleteLXC: %w", err)
}

fmt.Println(upid.Data) // e.g. UPID:pve1:...
```

Returns [`*types.DeleteLXCResponse`](/types/lxc.go).
> [!NOTE]
> Request time out after 30 seconds. A timeout returns a wrapped error. Check with `errors.Is(err, context.DeadlineExceeded)`.

**TypeScript**
```typescript
const upid = await client.node("pve1").lxc(2222).deleteLXC();
if (upid instanceof Error) {
    console.error("lxc.deleteLXC:", upid);
    return;
}

console.log(upid.Data); // e.g. UPID:pve1:...
```

> [!IMPORTANT]
> This request is asynchronous. Use the returned UPID to track task progress. See [Tasks documentation](../nodes/tasks.md).

> [!WARNING]
> The container must be stopped before deletion. Attempting to delete a running container returns an error.

---

## CloneLXC

Clones a container to a new VMID, optionally on a different node.  
**Proxmox API:** [`POST /nodes/{node}/lxc/{vmid}/clone`](https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/clone)

**Go**
```go
func (s *LXCService) CloneLXC(data types.CloneLXCData) (*types.CloneLXCResponse, error)
```
```go
upid, err := client.Node("pve1").LXC(2222).CloneLXC(types.CloneLXCData{
    // Mandatory
    NewId: 2223, // use client.Cluster().GetNextId() to get a free VMID

    // Optional
    Target: "pve2", // clone to a different node
})
if err != nil {
    return fmt.Errorf("lxc.CloneLXC: %w", err)
}

fmt.Println(upid.Data) // e.g. UPID:pve1:...
```

Returns [`*types.CloneLXCResponse`](/types/lxc.go).
> [!NOTE]
> Request time out after 30 seconds. A timeout returns a wrapped error. Check with `errors.Is(err, context.DeadlineExceeded)`.

**TypeScript**
```typescript
const upid = await client.node("pve1").lxc(2222).cloneLXC({
    // Mandatory
    NewId: 2223, // use client.cluster().getNextId() to get a free VMID

    // Optional
    Target: "pve2", // clone to a different node
});
if (upid instanceof Error) {
    console.error("lxc.cloneLXC:", upid);
    return;
}

console.log(upid.Data); // e.g. UPID:pve1:...
```

> [!IMPORTANT]
> This request is asynchronous. Use the returned UPID to track task progress. See [Tasks documentation](../nodes/tasks.md).

---

## See Also

- [Client documentation](/docs/getting-started/authentication.md)
- [Other Guides](/docs/guides/)
- [LXC associated types](/types/lxc.go)
- [Proxmox API - LXC endpoints](https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc)
- [Proxmox Wiki - Linux Container](https://pve.proxmox.com/wiki/Linux_Container)