# Changelog

All notable changes to this project will be documented in this file.
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Retryable client requesting Proxmox API.
- Version service: `GetVersion`
- Cluster service: `GetTasks`, `GetNextId`
- Nodes service: `GetNodes`
- Tasks service: `GetTasks`, `GetTaskStatus`, `DeleteTask`
- LXC service: `GetLXCs`, `PostLXC`, `DeleteLXC`, `CloneLXC`
- LXC Status service: `StartLXC`, `StopLXC`
- WASM: Go code interfacing for TypeScript consumption
- Tests: black-box testing
- CI/CD: GitHub Actions (lint, build, tests, publish packages, mirroring)
- CI/CD: Makefile (format & lint, build, tests, deps, cleanup, help)
- Documentation: Getting Started (installation, authentication, configuration)
- Documentation: Guides (cluster, nodes, tasks, lxc, lxc status, version)
- Documentation: Contributing (architecture, contributing, testing)
- Documentation: Reference (api-reference, errors, types)
- License: MIT, third-party notices for the npm package (WASM bundle)
- Tooling: `vm.sh` script to set up a Proxmox VE virtual machine (QEMU)
- GitHub: issue templates (bug, feature, documentation) and pull request template
- Documentation: Open source community files (Security policy, Code of Conduct)