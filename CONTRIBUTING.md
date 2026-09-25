# Contributing

> ProxmoxSDK is not open to direct code contributions at this time. However, external feedback, technical reviews, and suggestions on existing code are welcome and appreciated.

## Table of Contents

- [Development Setup](#development-setup)
- [Workflow](#workflow)
  - [Branches](#branches)
  - [Commits](#commits)
- [Code Conventions](#code-conventions)
- [Pull Requests](#pull-requests)
- [Reporting Bugs](#reporting-bugs)
- [Suggesting Features](#suggesting-features)
- [Running Tests](#running-tests)
- [See Also](#see-also)

---

## Development Setup

**Prerequisites**
- Go 1.26+
- Node.js 25+
- Git

**Clone the repository**
```bash
git clone https://github.com/HavelCTF/ProxmoxSDK.git
cd ProxmoxSDK
```

**Install dependencies**
```bash
make deps
```

**Build**
```bash
make build
```

---

## Workflow

### Branches

| Type | Pattern | Example |
|------|---------|---------|
| Issue branch | `issue-{issue-number}/{author}/{description}` | `issue-1/MaxenceLgt/resolve-doc-install-proxmox` |
| Dev branch | `dev` | `dev` |
| Stable | `main` | `main` |

Branches are created from `dev`. `main` receives merges from `dev` only on validated releases.

### Commits

Commits follow the [Karma](http://karma-runner.github.io/6.4/dev/git-commit-msg.html) convention with scope:
```
feat(scope): short description
^--^ ^---^   ^--------------^
|    |        |
|    |        └── imperative, lowercase, no period
|    └─────────── affected area (e.g. lxc, cluster, docs, ci)
└──────────────── type (see below)
```

**Types**

| Type | Usage |
|------|-------|
| `feat` | New feature |
| `fix` | Bug fix |
| `docs` | Documentation only |
| `refactor` | Code change that is not a fix nor a feature |
| `test` | Adding or updating tests |
| `chore` | Build, CI, dependencies |
| `ci` | GitHub Actions, Makefile |

---

## Code Conventions

### Go
Code follows the guidelines defined in [Effective Go](https://go.dev/doc/effective_go).

### TypeScript
Code follows the established TypeScript conventions:

- **Naming:** `camelCase` for variables and functions, `PascalCase` for classes and types. Well-known acronyms keep their capitalization (e.g. `LXCService`, `WASMLoader`), `UPPER_SNAKE_CASE` for constants
- **Types:** explicit type annotations on function signatures. Avoid `any`
- **Imports:** named imports preferred over default imports
- **Async:** `async/await` over raw Promises
- Linting and formatting enforced via [Biome](https://biomejs.dev/)

---

## Pull Requests

Pull requests use the **CPSE format** in their description:
```markdown
**Context:** General context of the repository and the area affected by this PR.

**Problem:** The specific problem this PR addresses.

**Solution:** Short bullet points describing the changes made.

**Expectation:** What is expected from reviewers: general opinion, full review, focus on a specific area, etc.
```

**Guidelines**
- One PR per feature or fix — keep scope focused
- Update documentation if the PR affects the public API
- All tests must pass before requesting a review (`make test`)

---

## Reporting Bugs

The project is currently under supervised access. If you encounter a bug, contact the development team directly via the channel through which you were given access to the project.

When reporting, include:
- Go version (`go version`)
- SDK version
- Steps to reproduce
- Expected behavior vs observed behavior

---

## Suggesting Features

Feature suggestions are not open until `v1.0.0` is released. The team is currently focused on delivering the mandatory features of the project. Suggestions will be welcome once the project reaches general availability.

---

## Running Tests
```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Run manually
go test -v ./...
```

For more details on the test suite and conventions, see [Testing](./docs/contributing/testing.md).

---

## See Also

- [Architecture](./architecture.md)
- [Testing](./docs/contributing/testing.md)
- [Effective Go](https://go.dev/doc/effective_go)
- [Karma Commit Convention](http://karma-runner.github.io/6.4/dev/git-commit-msg.html)