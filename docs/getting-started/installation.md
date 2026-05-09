# Installation

Guide to install Proxmox SDK in your project.

## Table of Contents
- [Go Installation](#go-installation)
  - [Prerequisites](#prerequisites)
  - [Installation Methods](#installation-methods)
- [TypeScript Installation](#typescript-installation)
  - [Prerequisites](#prerequisites-1)
  - [Installation Methods](#installation-methods-1)
- [Troubleshooting](#troubleshooting)
- [Next Steps](#next-steps)

## Go Installation

### Prerequisites

- **Go 1.26+**
- **Access to private repository**

### Installation Methods

#### Install via go get
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

#### Verify Installation
```bash
go list -m github.com/HavelCTF/ProxmoxSDK
```

**Expected output:**
```
github.com/HavelCTF/ProxmoxSDK
```

## TypeScript Installation

### Prerequisites

- **Node.js 25+ (LTS recommended)**
- **npm**

### Installation Methods

#### Install via npm
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

## Troubleshooting

**Problem:** "module not found"  
**Solution:** Run `go mod tidy`

**Problem:** "Go version too old"  
**Solution:** Update Go to 1.26+: `go version` to check current version

**Problem:** "npm package not found"  
**Solution:** Check package name is `@havelctf/proxmox-sdk`

## Next Steps

After successful installation:

1. **[Create your first client](authentication.md)** - Set up authentication and connect to Proxmox
2. **[Configure options](configuration.md)** - Customize client behavior
3. **[Read guides](../guides/)** - Learn to use specific services