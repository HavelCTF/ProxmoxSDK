# ProxmoxSDK Monorepo

A dual-implementation SDK for Proxmox VE in both **TypeScript** and **Go**, designed to provide identical functionality and behavior across both languages.

## 🏗️ Repository Structure

```
ProxmoxSDK/
├── cmd
│   └── wasm
│       └── main_wasm.go
├── go.mod
├── go.sum
├── internal
│   ├── client
│   │   ├── client.go
│   │   └── request.go
│   ├── context
│   │   └── lxc
│   │       └── context.go
│   ├── http
│   │   └── http.go
│   ├── lxc
│   │   ├── instance
│   │   │   └── instance.go
│   │   └── lxc.go
│   ├── nodes
│   │   └── nodes.go
│   └── version
│       └── version.go
├── Makefile
├── pkg
│   ├── proxmox
│   │   └── proxmox.go
│   └── typescript
│       └── example.ts
├── README.md
├── test
│   └── spec_runner_test.go
└── types
    ├── lxc.go
    ├── nodes.go
    └── version.go
```

## 🚀 Quick Start

### Prerequisites

- **Go** (v1.23+)

### Installation

```
Work in Progress
```

### Building

```
Work in Progress
```

### Testing

```
Work in Progress
```

## 📚 SDK Usage

### TypeScript

```
Work in Progress
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
        "https://proxmox.example.com",
        "PVEAPIToken=user@pam!token=...",
        "unique-client-id",
    )

    // Get version information
    version, err := client.Version()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Version: %s\n", version.Data)
}
```

## 🛠️ Development

### Code Formatting

```
Work in Progress
```

### Linting

```
Work in Progress
```

### Cleaning Build Artifacts

```
Work in Progress
```

## 📦 Publishing

### TypeScript Package

```
Work in Progress
```

### Go Module

The Go module is automatically available via:
```bash
go get github.com/HavelCTF/ProxmoxSDK
```

## 🎯 Design Goals

1. **Feature Parity**: Both implementations provide identical functionality
2. **Behavioral Consistency**: Same inputs produce same outputs across languages
3. **Idiomatic Code**: Each implementation follows language-specific best practices
4. **Maintainability**: Parallel development with automated behavioral verification
5. **Testing**: Comprehensive unit tests and cross-language behavioral tests

## 🤝 Contributing

When adding new features:

1. Implement the feature in both TypeScript and Go
2. Add unit tests in both implementations
4. Update documentation
5. Ensure all tests pass

## 📝 Available Commands

```
Work in Progress
```

## 📄 License

ISC

## 🔗 Links

- [Repository](https://github.com/HavelCTF/ProxmoxSDK)
- [Issues](https://github.com/HavelCTF/ProxmoxSDK/issues)
- [Proxmox VE API Documentation](https://pve.proxmox.com/pve-docs/api-viewer/)
