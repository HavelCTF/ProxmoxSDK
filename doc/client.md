# Client

> Client is the entry point of the SDK. It handles authenticated HTTP requests to the Proxmox API and exposes all available services.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Go](#go)
  - [Initialization](#initialization)
- [TypeScript](#typescript)
  - [Initialization](#initialization-1)
- [See Also](#see-also)

## Prerequisites

- A running Proxmox instance
- A Proxmox API token with sufficient permissions for the requests you intend to make
- A backend generated UUID for logging

## Go

### Initialization

```go
package main

import (
    "fmt"
    "github.com/HavelCTF/ProxmoxSDK/pkg/proxmox"
)

func main() {
    client := proxmox.NewClient(
        "https://proxmox.example.com", // Proxmox API base URL
        "PVEAPIToken=user@pam!tokenid=uuid", // Proxmox API token
        "unique-client-id", // Backend generated UUID
    )

    fmt.Println(client.UUID())
    fmt.Println(client.BaseURL())
}
```

## TypeScript

> Work In Progress

## See Also

<details>
<summary>Next</summary>

- [Nodes](./nodes/nodes.md)
- [Version](./version.md)
- [Cluster](./cluster/cluster.md)

</details>