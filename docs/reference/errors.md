# Errors

> All errors returned by ProxmoxSDK are standard Go errors wrapped with `fmt.Errorf`.

## Table of Contents

- [Error Format](#error-format)
- [Inspecting Errors](#inspecting-errors)
- [Known Error Cases](#known-error-cases)
- [See Also](#see-also)

---

## Error Format

Error messages vary depending on where the error occurred in the request pipeline.

**Request executed — non-OK status returned**
```
[<uuid>] (<baseURL><endpoint>) Request failed with status <status_code>: <body>
```
```
[test-0] (http://your-host:8006/api2/json/version) Request failed with status 500: internal server error
```

**Error before `DoRequest` — no HTTP call made**
```
[<uuid>] <custom message>: <error caught>
```
```
[test-0] failed to serialize request body: unexpected end of JSON input
```

| Field | Present | Description |
|-------|---------|-------------|
| `uuid` | Always | Client UUID passed at `NewClient` initialization |
| `baseURL` + `endpoint` | Request executed only | Full URL of the called endpoint |
| `status_code` | Non-OK response only | HTTP status code returned by Proxmox |
| `body` | Non-OK response only | Raw response body returned by Proxmox |
| `custom message` | Always | Description of what failed |
| `error caught` | Always | Underlying Go error |

---

## Inspecting Errors

**Check for a timeout**
```go
if errors.Is(err, context.DeadlineExceeded) {
    // request exceeded the 30s internal timeout
}
```

**Log the full error**
```go
version, err := client.GetVersion()
if err != nil {
    log.Println(err)
    // [test-0] (http://your-host:8006/api2/json/version) Request failed with status 500: internal server error
}
```

---

## Known Error Cases

| Method | Condition | Status |
|--------|-----------|--------|
| Any | Request exceeds 30s internal timeout | `context.DeadlineExceeded` |
| Any | Invalid or malformed API token | `401` |
| Any | Token does not have sufficient privileges | `403` |
| Any | Proxmox returns `5xx` after 3 retries | `5xx` |
| Any `POST` | Server unavailable — no retry | `5xx` |
| `Cluster.GetNextId` | Requested VMID already in use | `400` |
| `Node.PostLXC` | VMID conflict on creation | `400` |
| `Node.PostLXC` | Template not found | `500` |
| `LXC.DeleteLXC` | Container is still running | `400` |
| `LXC.CloneLXC` | Target VMID already in use | `400` |

> [!NOTE]
> `POST` requests are never retried. If a creation or action call fails due to a transient error, verify the resource state before retrying manually.

> [!WARNING]
> The internal timeout of 30 seconds is not configurable. If `context.DeadlineExceeded` is returned, verify that the Proxmox instance is reachable and responsive.

---

## See Also

- [API Reference](./api-reference.md)
- [Contributing — Architecture](../contributing/architecture.md)
- [Go errors package](https://pkg.go.dev/errors)
- [Go fmt.Errorf](https://pkg.go.dev/fmt#Errorf)