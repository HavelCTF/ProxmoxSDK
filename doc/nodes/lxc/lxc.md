# LXC

> Entry point for all lxc-scoped services from selected node (creation, deletion, ...).

## Table of Contents

- [Prerequisites](#prerequisites)
- [Go](#go)
  - [Service Initialization](#service-initialization)
  - [GetLXCs](#getlxcs)
  - [PostLXC](#postlxc)
  - [DeleteLXC](#deletelxc)
- [TypeScript](#typescript)
- [See Also](#see-also)

## Prerequisites

Requires an initialized node service. See the [Node documentation](../nodes.md).

## Go

### Service Initialization

Initializes an LXC service from an existing node service.
```go
lxcService := nodeService.LXC(vmid)

fmt.Println(lxcService.Node())
fmt.Println(lxcService.VMID())
```

### GetLXCs

Returns all LXCs information from a node.

```go
lxcs, err := nodeService.GetLXCs() // *types.LXCsResponse, error
```

See the [LXCsResponse](types/lxc.go) type for available fields.

### PostLXC

Creates a container using given information.

```go
import (
    "github.com/HavelCTF/ProxmoxSDK/types"
)

lxcData :=  types.CreateLXCData{
    // Mandatory
    Node: nodeService.Node(),
    OSTemplate: "path-to-template", // e.g : local:vztmpl/debian-12-standard_12.12-1_amd64.tar.zst
    VMID: 2222, // Use strconv.Atoi(nodeService.Cluster().GetNextId()) result

    // Optional
    Features: &types.LXCFeatures{
        Nesting: true,
    },

}

res, err := nodeService.PostLXC(lxcData) // *types.CreateLXCResponse, error
```

Creation request is asynchronous. Check task progress using the returned UPID. See [Node Documentation](../nodes.md).

See the [CreateLXCResponse](types/lxc.go) type for available fields.

### DeleteLXC

Deletes a container.

```go
res, err := lxcService.DeleteLXC() // *types.DeleteLXCResponse, error

// OR

res, err := nodeService.LXC(vmid).DeleteLXC() // *types.DeleteLXCResponse, error
```

Deletion request is asynchronous. Check task progress using the returned UPID. See [Node Documentation](../nodes.md).

See the [DeleteLXCResponse](types/lxc.go) type for available fields.

## TypeScript

> Work In Progress

## See Also

<details>
<summary>Previous</summary>

- [Nodes](../nodes.md)

</details>