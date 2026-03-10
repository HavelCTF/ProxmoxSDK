# Status

> Entry point for all lxc-status-scoped services from selected lxc (start, stop, ...).

## Table of Contents

- [Prerequisites](#prerequisites)
- [Go](#go)
  - [Service Initialization](#service-initialization)
  - [StartLXC](#startlxc)
  - [StopLXC](#stoplxc)
- [TypeScript](#typescript)
- [See Also](#see-also)

## Prerequisites

Requires an initialized lxc service. See the [LXC documentation](../lxc.md).

## Go

### Service Initialization

Initializes a Status service from an existing lxc service.
```go
lxcService := lxcService.Status()
```

### StartLXC

Starts a container.

```go
lxcs, err := statusService.StartLXC() // *types.StartLXCResponse, error
```

Start request is asynchronous. Check task progress using the returned UPID. See [Tasks Documentation](../../tasks/tasks.md).

See the [StartLXCResponse](/types/lxc.go) type for available fields.

### StopLXC

Stops a container.

```go
lxcs, err := statusService.StopLXC() // *types.StopLXCResponse, error
```

Stop request is asynchronous. Check task progress using the returned UPID. See [Tasks Documentation](../../tasks/tasks.md).

See the [StopLXCResponse](/types/lxc.go) type for available fields.

## TypeScript

> Work In Progress

## See Also

<details>
<summary>Previous</summary>

- [LXC](../lxc.md)

</details>
