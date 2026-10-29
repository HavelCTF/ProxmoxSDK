# Installation

Guide to install Proxmox SDK and its requirements in your project.

## Table of Contents
- [Proxmox VE](#proxmox-ve)
  - [Using the script (QEMU)](#using-the-script-qemu)
  - [From the website (manual)](#from-the-website-manual)
- [Go Installation](#go-installation)
  - [Prerequisites](#prerequisites)
  - [Installation Methods](#installation-methods)
- [TypeScript Installation](#typescript-installation)
  - [Prerequisites](#prerequisites-1)
  - [Installation Methods](#installation-methods-1)
- [Troubleshooting](#troubleshooting)
- [Next Steps](#next-steps)

## Proxmox VE

The SDK talks to a Proxmox VE server. If you don't have one, you can install Proxmox VE in a virtual machine, either with the script provided in the repository or manually.

### Using the script (QEMU)

The `vm.sh` script, at the root of the repository, downloads the Proxmox VE installer and runs it in a QEMU virtual machine.

**Prerequisites**

- A **Linux** machine with **KVM** enabled (`/dev/kvm` exists and your user is allowed to use it, usually through the `kvm` group)
- A **graphical session**: the virtual machine opens in a window
- About **50 GB** of free disk space, **4 GB** of RAM and **2 CPUs** (default values, editable at the top of `vm.sh`)
- QEMU. On Debian / Ubuntu:
  ```bash
  sudo apt install qemu-system-x86 qemu-utils
  ```

**1. Setup**

```bash
./vm.sh setup
```

Downloads the Proxmox VE ISO, verifies its checksum and creates the virtual disk `disk.qcow2`. The command can be run again safely: it only does what is missing, and an interrupted download is resumed.

**2. Install Proxmox VE**

```bash
./vm.sh install
```

Boots the ISO in a window. Follow the Proxmox VE installer (it asks for a password for the `root` user). When the installation is over, **close the window**: the virtual machine would otherwise boot the installer again.

**3. Run the virtual machine**

```bash
./vm.sh
```

Boots the installed system from the virtual disk. Once started, the virtual machine is available from your machine on:

| Service | Address |
|---------|---------|
| Proxmox web interface | `https://localhost:8006` |
| SSH | `ssh -p 2222 root@localhost` |

The web interface uses a self-signed certificate, so your browser shows a warning the first time. Log in as `root` with the password chosen during the installation (realm **Linux PAM**), then [create an API token](authentication.md#creating-an-api-token). Use `https://localhost:8006` as the URL of your client.

> [!NOTE]
> Closing the window stops the virtual machine. Its data is kept in `disk.qcow2`: to start from scratch, delete this file and run `./vm.sh setup` and `./vm.sh install` again.

### From the website (manual)

- Download the ISO installer from the [Proxmox website](https://www.proxmox.com/en/downloads/proxmox-virtual-environment).
- Install and run virtual machine software (VirtualBox, VMware, QEMU, ...).
- Create a virtual machine that meets the [hardware requirements](https://www.proxmox.com/en/products/proxmox-virtual-environment/requirements) and boot the ISO in it.
- Make sure the virtual machine is reachable from your machine over HTTPS on port `8006` (bridged network, or port forwarding).

## Go Installation
 
### Prerequisites
 
- **Go 1.26+**

### Installation Methods
 
**Install via go get**

```bash
go get github.com/HavelCTF/ProxmoxSDK
```

To install a specific version, branch or commit:

```bash
go get github.com/HavelCTF/ProxmoxSDK@dev      # a branch
go get github.com/HavelCTF/ProxmoxSDK@<commit> # a commit
```
 
**Verify Installation**

```bash
go list -m github.com/HavelCTF/ProxmoxSDK
```
 
**Expected output:** the module path followed by the installed version
```
github.com/HavelCTF/ProxmoxSDK v1.0.0
```

## TypeScript Installation

### Prerequisites

- **Node.js 25+ (LTS recommended)**
- **npm**

### Installation Methods

#### Install via npm
> [!NOTE]
> You need to have a Personal Access Token (PAT) on GitHub: `Settings > Developer Settings > PAT > Tokens (classic) > Generate new token (classic)` with the `read:packages` scope

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

**Problem:** "npm package not found", or an authentication error (401 / 404) from npm  
**Solution:** Check the package name is `@havelctf/proxmox-sdk`, that your `.npmrc` contains the `@havelctf` registry line, and that `NPM_TOKEN` is set with a token having the `read:packages` scope

**Problem:** `./vm.sh` fails with a KVM error (permission denied or no such device)  
**Solution:** Check that virtualization is enabled in your BIOS / UEFI, that `/dev/kvm` exists, and that your user belongs to the `kvm` group (`sudo usermod -aG kvm $USER`, then log in again)

**Problem:** `no disk.qcow2` or `no ISO` when running `./vm.sh`  
**Solution:** Run `./vm.sh setup` first

**Problem:** `https://localhost:8006` does not answer  
**Solution:** Check that the virtual machine is running (`./vm.sh`) and give it a minute to finish booting

## Next Steps

After successful installation:

1. **[Create your first client](authentication.md)** - Set up authentication and connect to Proxmox
2. **[Configure options](configuration.md)** - Customize client behavior
3. **[Read guides](../guides/)** - Learn to use specific services