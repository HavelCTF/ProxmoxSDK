# Configuration

Guide to configuring the Proxmox SDK client.

## Table of Contents
- [Default Configuration](#default-configuration)
- [Custom HTTP Client](#custom-http-client)
- [Examples](#examples)
- [Next Steps](#next-steps)

## Default Configuration

### Go

By default, the SDK uses an HTTP client with:

- **Timeout:** 60 seconds
- **TLS Verification:** Disabled (accepts self-signed certificates)

```go
// Default client (used automatically)
&http.Client{
    Transport: &http.Transport{
        TLSClientConfig: &tls.Config{
            InsecureSkipVerify: true, // ⚠️ Development only
        },
    },
    Timeout: 60 * time.Second,
}
```

> [!WARNING]
> The default client skips TLS verification to work with self-signed certificates commonly used in Proxmox installations. **For production**, use a custom client with proper certificate validation.

---

### TypeScript

By default, the SDK uses an HTTP client with:

- **Timeout:** 60 seconds
- **TLS Verification:** Enabled
- **Transport:** custom `http.RoundTripper` that routes HTTP requests through the JavaScript `fetch()` API, enabling network calls from a WASM binary running in Node.js or a browser.

```go
// Default client (used automatically)
return &http.Client{
    Transport: &fetchTransport{},
    Timeout:   60 * time.Second,
}
```

---

## Custom HTTP Client

### Go

Provide your own `http.Client` using `WithHTTPClient`:
```go
customClient := &http.Client{
    Timeout: 30 * time.Second,
}

client := proxmox.NewClient(
    "https://proxmox.example.com:8006",
    token,
    "my-app",
    proxmox.WithHTTPClient(customClient),
)
```

For production, enable TLS verification with a custom CA certificate:
```go
caCert, _ := os.ReadFile("/path/to/proxmox-ca.crt")
caCertPool := x509.NewCertPool()
caCertPool.AppendCertsFromPEM(caCert)

client := proxmox.NewClient(
    "https://proxmox.example.com:8006",
    token,
    "production-app",
    proxmox.WithHTTPClient(&http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                RootCAs: caCertPool,
            },
        },
        Timeout: 60 * time.Second,
    }),
)
```

---

### TypeScript

Use the `insecure` option to disable TLS verification for self-signed certificates:
```typescript
const client = await ProxmoxSDK.create(
    "https://proxmox.example.com:8006",
    token,
    "my-app",
    { insecure: true } // ⚠️ Development only
);
```

> [!WARNING]
> `insecure: true` sets `NODE_TLS_REJECT_UNAUTHORIZED=0` at the process level. This affects all HTTPS connections in the Node.js process. **Do not use in production.**

> [!NOTE]
> Custom HTTP client configuration is not yet exposed in the TypeScript SDK. For advanced TLS configuration, use the Go SDK directly. Full HTTP client customization is planned for a future release.

---

## Examples

### Go Production Client
```go
func newProductionClient(url, token string) *proxmox.Client {
    caCert, err := os.ReadFile("/etc/ssl/certs/proxmox-ca.crt")
    if err != nil {
        log.Fatal("Failed to load CA certificate:", err)
    }

    caCertPool := x509.NewCertPool()
    if !caCertPool.AppendCertsFromPEM(caCert) {
        log.Fatal("Failed to parse CA certificate")
    }

    return proxmox.NewClient(url, token, "production-app",
        proxmox.WithHTTPClient(&http.Client{
            Transport: &http.Transport{
                TLSClientConfig: &tls.Config{
                    RootCAs: caCertPool,
                },
            },
            Timeout: 60 * time.Second,
        }),
    )
}
```

### TypeScript Development Client
```typescript
const client = await ProxmoxSDK.create(
    "https://proxmox.example.com:8006",
    token,
    "dev-app",
    { insecure: true }
);
```

---

## Next Steps

After configuring your client:

1. **[Read service guides](../guides/)** - Learn to use specific features
2. **[API Reference](../reference/api-reference.md)** - Complete method documentation

## See Also

- [Authentication Guide](authentication.md)
- [Installation Guide](installation.md)
- [Go net/http documentation](https://pkg.go.dev/net/http)