# ProxmoxSDK
 
> Go + TypeScript SDK for the Proxmox VE API.
 
![Go Version](https://img.shields.io/badge/go-1.26+-blue)
![License](https://img.shields.io/badge/license-MIT-green)
[![codecov](https://codecov.havel-ctf.com/github/HavelCTF/ProxmoxSDK/graph/badge.svg?token=LLSE7FRBHK)](https://codecov.havel-ctf.com/github/HavelCTF/ProxmoxSDK)
 
> [!NOTE]
> **Project status:** under active development, no stable release has been published yet. The public API may still change. See the [Changelog](docs/CHANGELOG.md).

## Requirements
 
- **Go 1.26+** to use the Go package
- **Node.js 25+** and npm to use the TypeScript package
- A **Proxmox VE** instance reachable over HTTPS (port `8006` by default) and an **API token**, see [Authentication](docs/getting-started/authentication.md)

> [!NOTE]
> You don't have a **Proxmox VE instance**? Follow the [Proxmox VE installation guide](#proxmox-ve) to set one up. It is only needed to **try the SDK on a real server**, the test suite runs without one.

## Installation

### Proxmox VE

The SDK needs a Proxmox VE instance. If you don't have one, follow the [installation guide](docs/getting-started/installation.md) to set one up.

**Go**
```bash
go get github.com/HavelCTF/ProxmoxSDK
```

**TypeScript**
```bash
npm install @havelctf/proxmox-sdk
```

## Quick Start

### Proxmox

To run Proxmox Virtual Machine, if installed using [script](#installation):

```
./vm.sh
```

### Go

```go
package main

import (
    "fmt"
    "log"

    "github.com/HavelCTF/ProxmoxSDK/pkg/proxmox"
)

func main() {
    client := proxmox.NewClient(
        "https://your-host:8006",
        "PVEAPIToken=root@pam!token=your-token-uuid",
        "your-app-uuid",
    )

    version, err := client.GetVersion()
    if err != nil {
        log.Fatalf("failed to get version: %v", err)
    }

    fmt.Printf("Proxmox VE %s\n", version.Data.Version)
}
```

### TypeScript

```typescript
import { ProxmoxSDK } from "@havelctf/proxmox-sdk";

(async () => {
    const client = await ProxmoxSDK.create(
        "https://your-host:8006",
        "PVEAPIToken=root@pam!token=your-token-uuid",
        "your-app-uuid",
    );

    const version = await client.getVersion();
    if (version instanceof Error) {
        console.error("failed to get version:", version);
        return;
    }

    console.log("Proxmox VE", version.Data.Version);
})();
```

## Documentation
 
The full documentation is indexed in [docs/README.md](docs/README.md).
 
- **Getting started:** [Installation](docs/getting-started/installation.md), [Authentication](docs/getting-started/authentication.md), [Configuration](docs/getting-started/configuration.md)
- **Guides:** [Version, Cluster, ...](docs/guides/)
- **Reference:** [API, Errors, Types](docs/reference/)
- [Changelog](docs/CHANGELOG.md)

## Development

```bash
# Install dependencies
make deps

# Build
make build

# Run tests
make test

# Run CI checks
make ci
```

See [Contributing](docs/contributing/contributing.md) for full guidelines.