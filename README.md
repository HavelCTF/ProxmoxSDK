# ProxmoxSDK

> Go + TypeScript SDK for the Proxmox VE API.

![Go Version](https://img.shields.io/badge/go-1.26+-blue)
![License](https://img.shields.io/badge/license-MIT-green)
[![codecov](https://codecov.havel-ctf.com/github/HavelCTF/ProxmoxSDK/graph/badge.svg?token=LLSE7FRBHK)](https://codecov.havel-ctf.com/github/HavelCTF/ProxmoxSDK)

## Prerequisites

- Go 1.26+ : https://go.dev/doc/install  
- [Proxmox VE Setup](#installation)
- (Optional) Typescript : https://www.typescriptlang.org/download/

## Installation

### Proxmox VE

You can install Proxmox directly from the website or using an automated script.

**From the website**

- Download ISO Installer file from the [website](https://www.proxmox.com/en/downloads/proxmox-virtual-environment).
- Install and run a Virtual Machine software (VirtualBox, VMWare, qemu, ...).
- Create a Virtual Machine and run the ISO file in it. You can find the hardware requirements in the [proxmox documentation](https://www.proxmox.com/en/products/proxmox-virtual-environment/requirements).

**Using script**

```
sudo apt install qemu-system-x86 qemu-utils # install script dependencies
./vm.sh setup # download the ISO and Virtual Machine disk
./vm.sh install # boot the ISO to install Proxmox in Virtual Machine
```

### Other dependencies

**Go**
```bash
go get github.com/HavelCTF/ProxmoxSDK
```

**TypeScript**
```bash
npm install @havelctf/proxmox-sdk
```

## Quick Start

**Proxmox**

To run Proxmox Virtual Machine, if installed using [script](#installation):

```
./vm.sh
```

**Go**
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

**TypeScript**
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

- [Getting Started](docs/getting-started/installation.md)
- [Guides](docs/guides/)
- [API Reference](docs/reference/api-reference.md)
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
