# Version

> Version service exposes the Proxmox instance version.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Go](#go)
  - [GetVersion](#getversion)
- [TypeScript](#typescript)
  - [GetVersion](#getversion-1)
- [See Also](#see-also)

## Prerequisites

- Initialized client. See [Client](client.md)

## Go

### GetVersion
```go
package main

import (
    "fmt"
    "github.com/HavelCTF/ProxmoxSDK/pkg/proxmox"
)

func main() {
    // ... Client initialization

    version, err := client.GetVersion() // *types.VersionResponse, error
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(version)
}
```

> See [`VersionResponse`](../types/version.go) for the full list of available fields.

## TypeScript

### GetVersion

> Work In Progress

## See Also

<details>
<summary>Previous</summary>

- [Client](client.md)

</details>