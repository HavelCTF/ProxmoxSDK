# Version

> Exposes the Proxmox instance version.

## Table of Contents

- [Prerequisites](#prerequisites)
- [GetVersion](#getversion)
- [See Also](#see-also)

## Prerequisites

Requires an initialized client. See the [Client documentation](../getting-started/authentication.md).

## GetVersion

Returns the version of the Proxmox instance.

**Go**
```go
func (c *Client) GetVersion() (*types.VersionResponse, error) 
```
```go
version, err := client.GetVersion()
if err != nil {
    return fmt.Errorf("version.GetVersion: %w", err)
}

fmt.Println(version.Data.Version) // e.g. 8.1.4
```

Returns [`*types.VersionResponse`](/types/version.go).

**TypeScript**
```typescript
// Work In Progress
```

---

## See Also

- [Client documentation](../getting-started/authentication.md)
- [Other Guides](/docs/guides/)
- [Version associated types](/types/version.go)
- [Proxmox API - Version endpoint](https://pve.proxmox.com/pve-docs/api-viewer/#/version)