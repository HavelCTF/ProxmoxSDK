# ProxmoxSDK

> Go + TypeScript SDK for the Proxmox VE API.

![Go Version](https://img.shields.io/badge/go-1.26+-blue)
![License](https://img.shields.io/badge/license-ISC-green)


## Installation
> [!WARNING]
> Make sure to have access to the private repository

**Go**

```bash
export GOPRIVATE=github.com/HavelCTF/ProxmoxSDK
```
```bash
# Use SSH over HTTPS
git config --global url."git@github.com:".insteadOf "https://github.com/"
```
```bash
go get github.com/HavelCTF/ProxmoxSDK
```
```bash
# Reset to HTTPS
git config --global --remove-section url."git@github.com:"
```

**TypeScript**
> [!NOTE]
> You need to have a Personal Access Token (PAT) on GitHub: `Settings > Developer Settings > PAT > Tokens (classic) > Generate new token (classic)` with appropriate rights

```bash
export NPM_TOKEN=your_pat
```

Make sure to have an `.npmrc` file at the root of your repository with the following content:
```
engine-strict=true

@havelctf:registry=https://npm.pkg.github.com
//npm.pkg.github.com/:_authToken=${NPM_TOKEN}
```

```bash
npm install @havelctf/proxmox-sdk@latest
```

Versions accessible here: [pkg versions](https://github.com/HavelCTF/ProxmoxSDK/pkgs/npm/proxmox-sdk)

## Quick Start

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
