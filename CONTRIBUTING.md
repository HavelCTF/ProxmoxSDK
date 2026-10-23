# Contributing

## Table of Contents

- [Development Setup](#development-setup)
- [Workflow](#workflow)
  - [Branches](#branches)
  - [Branch Protection](#branch-protection)
  - [Commits](#commits)
- [Code Conventions](#code-conventions)
- [Pull Requests](#pull-requests)
- [Reporting Bugs](#reporting-bugs)
- [Suggesting Features](#suggesting-features)
- [Running Tests](#running-tests)
- [Licensing of contributions](#licensing-of-contributions)
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
| Issue branch | `{general-type}/({issue-number}-){description}` | `docs/(7-)resolve-doc-install-proxmox` |
| Dev branch | `dev` | `dev` |
| Stable | `main` | `main` |

Branches are created from `dev`. `main` receives merges from `dev` only on validated releases.

The `{general-type}` prefix is based on the [Karma](http://karma-runner.github.io/6.4/dev/git-commit-msg.html) types used for commits (see [Commits](#commits)). It describes the main purpose of the branch:

| General type | Usage | Example |
|--------------|-------|---------|
| `feat` | New feature or addition | `feat/12-add-storage-service` |
| `fix` | Bug fix | `fix/15-retry-on-post` |
| `docs` | Documentation only | `docs/7-resolve-doc-install-proxmox` |
| `refactor` | Code change that is not a fix nor a feature | `refactor/split-lxc-service` |
| `test` | Adding or updating tests | `test/task-wait-coverage` |
| `chore` | Build, CI, dependencies | `chore/bump-biome` |
| `ci` | GitHub Actions, Makefile | `ci/add-codecov-upload` |

- The `{issue-number}-` part is optional: omit it when the branch is not linked to an issue.
- `{description}` is short, lowercase and uses hyphens as separators.
- Pick the type matching the **main** change of the branch. The squash commit message on merge should use the same type.

### Branch Protection

`main` is protected by a [repository ruleset](https://github.com/HavelCTF/ProxmoxSDK/rules/25721394) (enforcement: `active`). The rule has **no bypass actor**: it applies to everyone, administrators included. Direct pushes to `main` are refused, and every change must go through a pull request.

| Rule | Requirement |
|------|-------------|
| Pull request | Changes reach `main` only through a pull request |
| Review | At least **1 approving review** from another member, never the author of the last push |
| Code owners | The approval must come from a code owner (`@HavelCTF/developers`, see [CODEOWNERS](./.github/CODEOWNERS)) |
| Stale reviews | New commits dismiss existing approvals, and the most recent push must be approved |
| Conversations | All review threads must be resolved before merging |
| Status checks | `lint-go`, `lint-typescript`, `check-liscence`, `test-go` and `build` must pass |
| Up to date | The branch must be up to date with `main` before merging |
| Signed commits | All commits must have a verified signature |
| Merge method | `squash` only |
| History | Force pushes and branch deletion are blocked |

> [!NOTE]
> A direct push is rejected with `GH013: Repository rule violations found`.

> [!IMPORTANT]
> Run `make ci` locally before opening a pull request: it runs the same lint, test and build steps as the required checks.

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

## Licensing of contributions

By submitting a contribution, you agree that it is licensed under the [MIT License](LICENSE) and that you have the right to submit it.
We use the Developer Certificate of Origin (https://developercertificate.org): sign your commits with `git commit -s`.

## See Also

- [Architecture](./architecture.md)
- [Testing](./docs/contributing/testing.md)
- [Effective Go](https://go.dev/doc/effective_go)
- [Karma Commit Convention](http://karma-runner.github.io/6.4/dev/git-commit-msg.html)