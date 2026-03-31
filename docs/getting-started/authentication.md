# Authentication

Complete guide to authenticate with Proxmox VE and create your first client.

## Table of Contents
- [Creating an API Token](#creating-an-api-token)
- [Token Format](#token-format)
- [Creating Your First Client](#creating-your-first-client)
  - [Go](#go)
  - [TypeScript](#typescript)
- [Permissions](#permissions)
- [Troubleshooting](#troubleshooting)
- [Next Steps](#next-steps)

## Creating an API Token

1. In your proxmox server, go to: **Datacenter** -> **Permissions** -> **API Tokens**

2. Click on **Add** and configure your token with following options:
    - **User**: Select the user (e.g., `root@pam`)
    - **Token ID**: Choose a descriptive identifier (e.g., `my-app`, `monitoring`, `backup`)
    - **Privilege Separation**: 
        - ⚠️ **Uncheck** for full user permissions (recommended for getting started)
        - ✅ **Check** for restricted permissions (recommended for production)
    - **Comment**: Optional description of token purpose
    - **Expire**: Optional expiration date (leave empty for no expiration)

> [!WARNING]
> The token secret is shown **only once**. So copy and save it immediately.

## Token Format

Proxmox API tokens follow this format:
```
PVEAPIToken=USERNAME@REALM!TOKENID=SECRET
```

**Components:**

| Component | Description | Example |
|-----------|-------------|---------|
| `USERNAME` | User account name | `root`, `admin`, `myuser` |
| `REALM` | Authentication realm | `pam` (Linux), `pve` (Proxmox) |
| `TOKENID` | Your chosen token identifier | `my-app`, `monitoring` |
| `SECRET` | The UUID secret from step 3 | `a1b2c3d4-e5f6-...` |

**Complete example:**
```
PVEAPIToken=root@pam!my-app=a1b2c3d4-e5f6-7890-abcd-ef1234567890
```

**Common realms:**
- `pam` - Linux PAM (most common)
- `pve` - Proxmox VE authentication
- `ldap` - LDAP/Active Directory (if configured)

## Creating Your First Client

### Go
```go
package main

import (
    "context"
    "fmt"
    "log"
    
    "github.com/HavelCTF/ProxmoxSDK/pkg/proxmox"
)

func main() {
    // Create client with your token
    client := proxmox.NewClient(
        "https://proxmox.example.com:8006", // Base URL
        "PVEAPIToken=root@pam!my-app=a1b2c3d4-e5f6-7890-abcd-ef1234567890", // Token
        "my-uuid", // App genereated uuid
    )
    
    // Test the connection
    version, err := client.GetVersion()
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    fmt.Printf("✅ Successfully connected to Proxmox VE %s\n", version.Data.Version)
    
    // List nodes to verify permissions
    nodes, err := client.GetNodes()
    if err != nil {
        log.Fatalf("Failed to list nodes: %v", err)
    }
    fmt.Printf("Found %d node(s):\n", len(nodes.Data))
    for _, node := range nodes.Data {
        fmt.Printf("  - %s (status: %s)\n", node.Node, node.Status)
    }
}
```

**Run:**
```bash
go run main.go
```

**Expected output:**
```
✅ Successfully connected to Proxmox VE 8.1.3
Found 1 node(s):
  - pve (status: online)
```

### TypeScript

```typescript
import { ProxmoxSDK } from "@havelctf/proxmox-sdk";

(async () => {
    // Create client with your token
    const client = await ProxmoxSDK.create(
        "https://your-host:8006", // Base URL
        "PVEAPIToken=root@pam!token=your-token-uuid", // Token
        "your-app-uuid", // App generated uuid
    );

    // Test the connection
    const version = await client.getVersion();
    if (version instanceof Error) {
        console.error("Failed to get version:", version);
        return;
    }

    console.log("✅ Successfully connected to Proxmox VE:", version.Data.Version)

    // List nodes to verify permissions
    const nodes = await client.getNodes()
    if (nodes instanceof Error) {
        console.error("Failed to get nodes:", nodes)
        return;
    }
    console.log(`Found ${nodes.Data.length} node(s):`);
    for (const node of nodes.Data) {
        console.log(`  - ${node.Node} (status: ${node.Status})`);
    }
})();
```

**Run:**
```bash
npx tsx index.ts
```

**Expected output:**
```
✅ Successfully connected to Proxmox VE: 8.1.3
Found 1 node(s):
  - pve (status: online)
```

## Permissions

### Default Permissions

If **Privilege Separation is disabled**, the token has **all permissions** of the user.

### Custom Permissions

If **Privilege Separation is enabled**, you must grant specific permissions.

**Common permissions for SDK usage:**

| Permission | Purpose |
|------------|---------|
| `VM.Audit` | Read VM/container information |
| `VM.Allocate` | Create VMs/containers |
| `VM.Config.*` | Modify VM/container configuration |
| `VM.PowerMgmt` | Start, stop, restart operations |
| `Datastore.Audit` | Read storage information |
| `Datastore.Allocate` | Create storage |

**Grant permissions via UI:**
1. Navigate to **Datacenter** → **Permissions**
2. Click **Add** → **API Token Permission**
3. Select your token & choose role (e.g., `PVEAdmin`, `VM.Audit`)

## Troubleshooting

### "401 Unauthorized"

**Cause:** Invalid or expired token

**Solutions:**
1. Verify token format includes `PVEAPIToken=` prefix
2. Check token hasn't been deleted in Proxmox
3. Ensure no extra spaces or newlines in token string
4. Try creating a new token

**Test token manually:**
```bash
curl -k -H "Authorization: PVEAPIToken=root@pam!my-app=xxx" \
  https://proxmox.example.com:8006/api2/json/version
```

### "403 Forbidden" / "Permission denied"

**Cause:** Token lacks required permissions

**Solutions:**
1. **If Privilege Separation is enabled:**
   - Grant necessary permissions (see [Permissions](#permissions))
   - Or disable Privilege Separation for testing

2. **Check user permissions:**
   - Ensure the user (e.g., `root@pam`) has the required permissions
   - User permissions cascade to token if Privilege Separation is disabled

3. **Verify ACLs:**
```bash
   pveum acl list
```

### "Connection refused" / "Network error"

**Cause:** Cannot reach Proxmox server

**Solutions:**
1. **Verify URL format:**
   - Must include `https://` (not `http://`)
   - Must include port `:8006`
   - Example: `https://proxmox.example.com:8006`

2. **Check network connectivity:**
```bash
   ping proxmox.example.com
   curl -k https://proxmox.example.com:8006
```

3. **Verify firewall allows port 8006:**
```bash
   # On Proxmox server
   ufw status
   iptables -L
```

4. **Check Proxmox is running:**
```bash
   systemctl status pveproxy
```

### Token format issues

**Common mistakes:**
```bash
❌ root@pam!my-app=xxx                    # Missing PVEAPIToken= prefix
❌ PVEAPIToken=root@pam:my-app=xxx        # Wrong separator (: instead of !)
❌ PVEAPIToken=root!my-app=xxx            # Missing realm (@pam)
✅ PVEAPIToken=root@pam!my-app=xxx        # Correct format
```

## Next Steps

After successful authentication:

1. **[Configure client options](configuration.md)** - Customize HTTP client
2. **[Read service guides](../guides/)** - Learn to use specific features
3. **[Explore reference](../reference/api-reference.md)** - Complete API documentation

## See Also

- [Installation Guide](installation.md)
- [Configuration Guide](configuration.md)
- [Proxmox API Token Documentation](https://pve.proxmox.com/pve-docs/pveum.1.html)
- [Proxmox Permission Management](https://pve.proxmox.com/wiki/User_Management)