# Nodes

> Entry point for all node-scoped services (LXC, Tasks, ...).

## Table of Contents

- [Prerequisites](#prerequisites)
- [Go](#go)
  - [Service Initialization](#service-initialization)
  - [GetNodes](#getnodes)
  - [GetNode](#getnode)
- [TypeScript](#typescript)
- [See Also](#see-also)

## Prerequisites

Requires an initialized client. See the [Client documentation](../client.md).

## Go

### Service Initialization

Initializes a node service from an existing client.
```go
nodeService := client.Node("node-name")
```

### GetNodes

Returns all nodes in the cluster.
```go
nodes, err := client.GetNodes() // *types.NodesResponse, error
```

See the [NodesResponse](types/nodes.go) type for available fields.

### GetNode

Returns detailed information about a specific node.
```go
node, err := nodeService.GetNode() // *types.NodeResponse, error

// OR

node, err := client.Node("node-name").GetNode() // *types.NodeResponse, error
```

See the [NodeResponse](types/nodes.go) type for available fields.

## TypeScript

> Work In Progress

## See Also

<details>
<summary>Previous</summary>

- [Client](../client.md)

</details>

<details>
<summary>Next</summary>

- [LXC](./lxc/lxc.md)
- [Tasks](./tasks/tasks.md)

</details>