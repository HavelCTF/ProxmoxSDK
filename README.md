# ProxmoxSDK Monorepo

A dual-implementation SDK for Proxmox VE in both **TypeScript** and **Go**, designed to provide identical functionality and behavior across both languages.

## 🏗️ Repository Structure

```
ProxmoxSDK/
├── typescript/          # TypeScript SDK implementation
│   ├── src/
│   │   ├── client.ts
│   │   ├── client.spec-test.ts
│   │   ├── index.ts
│   │   └── version/
│   │       ├── version.ts
│   │       └── version.spec-test.ts
│   ├── package.json
│   ├── tsconfig.json
│   └── biome.json
├── go/                  # Go SDK implementation
│   ├── client.go
│   ├── client_test.go
│   ├── client_spec_test.go
│   ├── version.go
│   ├── version_test.go
│   ├── version_spec_test.go
│   ├── spec_test_runner.go
│   ├── go.mod
│   └── go.sum
├── test-specs/          # Shared test specifications
│   ├── client-spec.yaml
│   ├── version-spec.yaml
│   └── README.md
├── build.py             # Unified build script
└── README.md            # This file
```

## 🚀 Quick Start

### Prerequisites

- **Python 3.7+** (for build script)
- **Node.js** (v18+) and npm
- **Go** (v1.21+)

### Installation

```bash
# Install all dependencies
./build.py install
```

### Building

```bash
# Build both SDKs
./build.py build

# Build individually
./build.py build --ts    # TypeScript only
./build.py build --go    # Go only
```

### Testing

```bash
# Run all tests (TS, Go, and spec-based tests)
./build.py test

# Run individual test suites
./build.py test --ts     # TypeScript tests
./build.py test --go     # Go tests
```

## 📚 SDK Usage

### TypeScript

```typescript
import { ProxmoxClient } from 'proxmoxsdk';
import 'proxmoxsdk/version';

const client = new ProxmoxClient(
  'https://proxmox.example.com',
  'PVEAPIToken=user@pam!token=...',
  'unique-client-id'
);

// Get version information
const version = await client.version();
console.log(version);
```

### Go

```go
package main

import (
    "fmt"
    "log"

    proxmox "github.com/HavelCTF/ProxmoxSDK/go"
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

    fmt.Printf("Version: %s\n", version.Version)
}
```

## 🧪 Shared Test Specifications

This monorepo uses a unique testing approach where tests are written once in YAML format and consumed by both TypeScript and Go test runners:

- **Test Specifications**: Located in `test-specs/`, define test scenarios including inputs, mock responses, and expected outputs
- **Language-Agnostic**: Both SDKs read the same YAML specs to ensure identical behavior
- **Mocking Support**: Each SDK implements mocking based on spec definitions
- **Behavioral Consistency**: Guarantees both implementations behave identically

Example test spec:
```yaml
tests:
  - name: "version_success"
    scenario:
      input:
        baseURL: "https://proxmox.example.com"
        apiToken: "test-token"
        uuid: "test-uuid"
      mock:
        endpoint: "/api2/json/version"
        statusCode: 200
        body:
          version: "8.0.3"
      expected:
        success: true
        output:
          version: "8.0.3"
```

See [test-specs/README.md](test-specs/README.md) for more details.

## 🛠️ Development

### Code Formatting

```bash
# Format all code
./build.py format

# Format individually
./build.py format --ts    # TypeScript
./build.py format --go    # Go
```

### Linting

```bash
# Lint all code
./build.py lint

# Lint individually
./build.py lint --ts    # TypeScript
./build.py lint --go    # Go
```

### Cleaning Build Artifacts

```bash
# Clean all
./build.py clean

# Clean individually
./build.py clean --ts    # TypeScript
./build.py clean --go    # Go
```

## 📦 Publishing

### TypeScript Package

```bash
cd typescript
npm version patch  # or minor, major
npm publish
```

### Go Module

The Go module is automatically available via:
```bash
go get github.com/HavelCTF/ProxmoxSDK/go
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
3. Add shared test specifications in `test-specs/` to verify behavioral consistency
4. Update documentation
5. Ensure all tests pass: `./build.py test`

## 📝 Available Commands

All commands use the `build.py` script:

- `./build.py install` - Install all dependencies
- `./build.py build [--ts|--go]` - Build SDKs
- `./build.py test [--ts|--go]` - Run tests
- `./build.py lint [--ts|--go]` - Lint code
- `./build.py format [--ts|--go]` - Format code
- `./build.py clean [--ts|--go]` - Clean build artifacts

Use `--ts` or `--go` flags to target specific SDK, otherwise both are processed.

## 📄 License

ISC

## 🔗 Links

- [Repository](https://github.com/HavelCTF/ProxmoxSDK)
- [Issues](https://github.com/HavelCTF/ProxmoxSDK/issues)
- [Proxmox VE API Documentation](https://pve.proxmox.com/pve-docs/api-viewer/)
