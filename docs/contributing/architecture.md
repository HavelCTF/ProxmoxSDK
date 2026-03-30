# Architecture

> Overview of the ProxmoxSDK internal structure and design decisions.

## Table of Contents

- [Repository Structure](#repository-structure)
- [Package Responsibilities](#package-responsibilities)
- [Client](#client)
  - [Initialization](#initialization)
  - [Retry Policy](#retry-policy)
  - [Request Pipeline](#request-pipeline)
- [Service Chain](#service-chain)
- [WASM](#wasm)
- [Types](#types)
- [See Also](#see-also)

---

## Repository Structure
```
.
├── cmd/
│   └── wasm/               # WASM entry point
├── internal/
│   ├── client/             # HTTP client, retry, transport
│   ├── cluster/            # Cluster service implementation
│   ├── http/               # DoRequest generic helper
│   ├── nodes/              # Nodes, Tasks, LXC, Status implementations
│   └── version/            # Version service implementation
├── pkg/
│   ├── proxmox/            # Public API — exported client and services
│   └── typescript/         # TypeScript package consuming the WASM build
├── types/                  # Shared request/response types
├── tests/                  # Black-box tests and mocks
└── docs/                   # Documentation
```

This layout follows the established Go project conventions:

- `internal/` — implementation details, never imported outside the module
- `pkg/` — public API surface exposed to consumers
- `cmd/` — executable entry points
- `types/` — shared data structures accessible to both `internal/` and `pkg/`

---

## Package Responsibilities

**Go**

| Package | Responsibility |
|---------|---------------|
| `pkg/proxmox` | Exported `Client`, service accessors, public API |
| `internal/client` | HTTP client construction, retry logic, transport |
| `internal/http` | `DoRequest`: generic typed request helper |
| `internal/cluster` | Cluster service internal implementation |
| `internal/nodes` | Node scoped services internal implementations (LXC, Tasks, ...) |
| `internal/version` | Version service internal implementation |
| `types/` | Shared request/response structs |
| `cmd/wasm` | WASM build entry point |

**TypeScript**

| Package | Responsibility |
|---------|---------------|
| Work In Progress | Work In Progress |

---

## Client

### Initialization

The public client is initialized via `NewClient` in `pkg/proxmox`:

**Go**
```go
func NewClient(baseURL string, token string, uuid string, opts ...ClientOption) *Client
```

**TypeScript**
```ts
// Work In Progress

// When completed, this section should document:
// - TypeScript equivalent of `NewClient` and its options
```

| Parameter | Description |
|-----------|-------------|
| `baseURL` | Proxmox instance URL (e.g. `https://your-host:8006`) |
| `token` | Proxmox API token (`PVEAPIToken=user@realm!name=uuid`) |
| `uuid` | Your App UUID used for logs |
| `opts` | Optional configuration via `ClientOption` |

**Available Go options**

| Option | Description |
|--------|-------------|
| `WithHTTPClient(c *http.Client)` | Override the HTTP client used by the retry client |

All services are available immediately after `NewClient` — no deferred initialization.

### Retry Policy

The retry policy is implemented at the Go level and applies transparently to both Go and TypeScript consumers via WASM.
```go
httpClient.RetryMax     = 3
httpClient.RetryWaitMin = 5 * time.Second
httpClient.RetryWaitMax = 5 * time.Second
```

| Rule | Detail |
|------|--------|
| Retried requests | `GET` only, server errors (`5xx`) and no-response |
| Never retried | `POST` requests, prevents duplicate resource creation |
| Max attempts | 3 |
| Wait between retries | Fixed 5 seconds |

> [!WARNING]
> `POST` requests are never retried. If a creation call fails, the caller is responsible for verifying the state before retrying manually.

### Request Pipeline

All service methods route through `internal/http.DoRequest`, a generic typed helper that handles serialization, authentication headers, and error mapping:
```go
func DoRequest[R any](ctx context.Context, c *client.Client, content RequestContent) (*R, error)
```

No HTTP primitives are exposed to the public API. Consumers interact only with Proxmox service methods and dedicated types.

---

## Service Chain

**Go**

Services are accessed through the public `Client` via accessor methods. Each accessor instantiates an internal service that holds a private reference to the underlying `client.Client`:
```go
// pkg/proxmox/client.go
func (c *Client) Node(node string) *nodes.NodeService {
    return nodes.New(c.c, node)
}
```

Sub-services are chained through context structs passed down the hierarchy:
```
client.Node("pve1")
└── nodes.New(client, "pve1") → NodeService
    └── .LXC(100)
        └── lxc.New(LXCContext{client, node, vmid}) → LXCService
            └── .Status()
                └── status.New(StatusContext{client, node, vmid}) → StatusService
```

The internal `client.Client` is never exposed. Each service layer receives only the context it needs and keeps it private.

**TypeScript**
```ts
// Work In Progress

// When completed, this section should document:
// - How the service chain is exposed through the WASM interface
```

---

## WASM
```ts
// Work In Progress

// When completed, this section should document:
// - How Go functions are exported and made callable from TypeScript
// - How `cmd/wasm/main_wasm.go` registers the exported functions
// - How to build the WASM binary and where the output is expected by the TypeScript package
```

The WASM entry point is located at `cmd/wasm/main_wasm.go`. The compiled output is consumed by the TypeScript package at `pkg/typescript/`, which exposes the SDK functionality to TypeScript consumers via `pkg/typescript/src/wasm/loader.ts`.

Platform-specific HTTP transport is handled transparently:

| File | Target |
|------|--------|
| `internal/client/transport_default.go` | Native Go builds |
| `internal/client/transport_wasm.go` | WASM builds (`GOARCH=wasm`) |
| `internal/client/roundtrip_wasm.go` | WASM HTTP round-tripper |

---

## Types

**Go**

All request and response types are defined in the `types/` package, organized by service:

| File | Contents |
|------|----------|
| `types/cluster.go` | `ClusterTasksResponse`, `ClusterNextIdResponse` |
| `types/nodes.go` | `NodesResponse`, `NodeTasksResponse`, `NodeTaskStatusResponse`, ... |
| `types/lxc.go` | `LXCsResponse`, `CreateLXCResponse`, `CloneLXCResponse`, ... |
| `types/version.go` | `VersionResponse` |

Types are accessible to both `internal/` and `pkg/` packages and are part of the public API surface.

**TypeScript**
```ts
// Work In Progress
```

---

## See Also

- [Contributing](./contributing.md)
- [Testing](./testing.md)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go project layout conventions](https://github.com/golang-standards/project-layout)