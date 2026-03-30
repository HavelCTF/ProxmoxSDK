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

### Installation Methods

#### Install via go get
```bash
go get github.com/HavelCTF/ProxmoxSDK
```

#### Install Specific Version
```bash
go get github.com/HavelCTF/ProxmoxSDK@v1.0.0
```

#### Update to Latest Version
```bash
go get -u github.com/HavelCTF/ProxmoxSDK
```

#### Verify Installation
```bash
go list -m github.com/HavelCTF/ProxmoxSDK
```

**Expected output:**
```
github.com/HavelCTF/ProxmoxSDK v1.0.0
```

#### Install via Docker
```dockerfile
FROM golang:1.26-alpine
RUN go get github.com/HavelCTF/ProxmoxSDK
```

## TypeScript Installation

### Prerequisites

- **Node.js 18+ (LTS recommended)**
- **npm or yarn**

### Installation Methods

#### Install via npm
```bash
npm install @havelctf/proxmox-sdk
```

#### Install via yarn
```bash
yarn add @havelctf/proxmox-sdk
```

#### Verify Installation
```bash
npm list @havelctf/proxmox-sdk
```

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