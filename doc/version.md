# Version

> Version service exposes the Proxmox instance version.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Go](#go)
  - [GetVersion](#getversion)
- [TypeScript](#typescript)
- [See Also](#see-also)

## Prerequisites

Requires an initialized client. See the [Client documentation](./client.md).


## Go

### GetVersion

Returns Proxmox instance version.

```go
version, err := client.GetVersion() // *types.VersionResponse, error
```

See the [VersionResponse](/types/version.go) type for available fields.

## TypeScript

> Work In Progress

## See Also

<details>
<summary>Previous</summary>

- [Client](./client.md)

</details>