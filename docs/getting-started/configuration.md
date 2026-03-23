# Configuration

Guide to configuring the Proxmox SDK client.

## Table of Contents
- [Default Configuration](#default-configuration)
- [Custom HTTP Client](#custom-http-client)
  - [Production Setup](#production-setup)
- [Examples](#examples)
- [Next Steps](#next-steps)

## Default Configuration

By default, the SDK uses an HTTP client with:

- **Timeout:** 60 seconds
- **TLS Verification:** Disabled (accepts self-signed certificates)
```go
// Default client (used automatically)
&http.Client{
    Transport: &http.Transport{
        TLSClientConfig: &tls.Config{
            InsecureSkipVerify: true,  // ⚠️ Development only
        },
    },
    Timeout: 60 * time.Second,
}
```

> [!WARNING]
> The default client skips TLS verification to work with self-signed certificates commonly used in Proxmox installations. **For production**, use a custom client with proper certificate validation.

---

## Custom HTTP Client

### Basic Usage

Provide your own `http.Client` using `WithHTTPClient`:
```go
import (
    "net/http"
    "time"
    "github.com/HavelCTF/ProxmoxSDK/pkg/proxmox"
)

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

---

### Production Setup

For production environments, enable TLS verification with a custom CA certificate:
```go
import (
    "crypto/tls"
    "crypto/x509"
    "net/http"
    "os"
    "time"
)

// Load custom CA certificate
caCert, err := os.ReadFile("/path/to/proxmox-ca.crt")
if err != nil {
    log.Fatal(err)
}

caCertPool := x509.NewCertPool()
caCertPool.AppendCertsFromPEM(caCert)

// Create production HTTP client
prodClient := &http.Client{
    Transport: &http.Transport{
        TLSClientConfig: &tls.Config{
            RootCAs: caCertPool,  // ✅ Verify certificates
        },
    },
    Timeout: 60 * time.Second,
}

client := proxmox.NewClient(
    "https://proxmox.example.com:8006",
    token,
    "production-app",
    proxmox.WithHTTPClient(prodClient),
)
```

---

## Examples

### Production Client
```go
import (
    "crypto/tls"
    "crypto/x509"
    "net/http"
    "os"
)

func newProductionClient(url, token string) *proxmox.Client {
    // Load CA certificate
    caCert, err := os.ReadFile("/etc/ssl/certs/proxmox-ca.crt")
    if err != nil {
        log.Fatal("Failed to load CA certificate:", err)
    }
    
    caCertPool := x509.NewCertPool()
    if !caCertPool.AppendCertsFromPEM(caCert) {
        log.Fatal("Failed to parse CA certificate")
    }
    
    // Create HTTP client with TLS verification
    httpClient := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                RootCAs: caCertPool,
            },
        },
        Timeout: 60 * time.Second,
    }
    
    return proxmox.NewClient(url, token, "production-app",
        proxmox.WithHTTPClient(httpClient),
    )
}
```

---

## Next Steps

After configuring your client:

1. **[Read service guides](../guides/)** - Learn to use specific features:
2. **[API Reference](../reference/api-reference.md)** - Complete method documentation

## See Also

- [Authentication Guide](authentication.md)
- [Installation Guide](installation.md)
- [Go net/http documentation](https://pkg.go.dev/net/http)