# Contributing

## Table of Contents

- [Development Setup](#development-setup)
- [Workflow](#workflow)
  - [Branches](#branches)
  - [Branch Protection](#branch-protection)
  - [Commits](#commits)
- [Code Conventions](#code-conventions)
  - [Go](#go)
  - [TypeScript](#typescript)
  - [License Headers](#license-headers)
- [Pull Requests](#pull-requests)
- [Issues](#issues)
  - [Reporting Bugs](#reporting-bugs)
  - [Suggesting Features](#suggesting-features)
  - [Documentation Issues](#documentation-issues)
  - [Security Vulnerabilities](#security-vulnerabilities)
- [Running Tests](#running-tests)
- [Licensing of contributions](#licensing-of-contributions)
- [See Also](#see-also)

---

## Development Setup

**Prerequisites**
- Go 1.26+
- Node.js 22+ (the version used by the CI)
- Git
- Make

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

Run `make help` to list every available target.

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
> Run `make ci` locally before opening a pull request: it runs the Go and TypeScript linters, the Go format check, the tests and the full build. The license header check (`check-liscence`) is not part of `make ci`, see [License Headers](#license-headers).

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

> [!IMPORTANT]
> **Every commit must be signed.** A commit without a verified signature is not accepted: it is required by the [branch protection](#branch-protection) rules, and signing off is also part of [Licensing of contributions](#licensing-of-contributions).
>
> Two different things are expected, and the same command covers both:
> - `-S` adds a **cryptographic signature** (GPG or SSH key registered on your GitHub account) so the commit shows as *Verified*.
> - `-s` adds the **DCO sign-off** (`Signed-off-by:` line).
>
> ```bash
> git commit -S -s -m "feat(lxc): add snapshot support"
> ```
>
> To sign every commit automatically: `git config --global commit.gpgsign true` (works for SSH keys too, with `gpg.format ssh`). See GitHub's guide on [signing commits](https://docs.github.com/en/authentication/managing-commit-signature-verification/signing-commits).

---

## Code Conventions

### Go
Code follows the guidelines defined in [Effective Go](https://go.dev/doc/effective_go).

- Formatting: `gofmt`. Format with `make fmt`, verify with `make fmt-check`
- Static analysis: `go vet`, run by `make lint-go`

### TypeScript
Code follows the established TypeScript conventions:

- **Naming:** `camelCase` for variables and functions, `PascalCase` for classes and types. Well-known acronyms keep their capitalization (e.g. `LXCService`, `WASMLoader`), `UPPER_SNAKE_CASE` for constants
- **Types:** explicit type annotations on function signatures. Avoid `any`
- **Imports:** named imports preferred over default imports
- **Async:** `async/await` over raw Promises
- Linting and formatting enforced via [Biome](https://biomejs.dev/): check with `make lint-ts`, auto-fix with `npm run lint:fix` from `pkg/typescript`

### License Headers

Every source file must start with the project's MIT license header:

```go
// SPDX-License-Identifier: MIT
// Copyright (c) 2026 HavelCTF
```

The `check-liscence` job fails when a file is missing it. To check or add headers locally, use [addlicense](https://github.com/google/addlicense):

```bash
go install github.com/google/addlicense@latest

# Check only (what the CI does)
addlicense -check -l mit -s=only -c "HavelCTF" -ignore="**/*.yml" .

# Add the missing headers
addlicense -l mit -s=only -c "HavelCTF" -ignore="**/*.yml" .
```

---

## Pull Requests

Pull requests are opened against `dev` and must follow the [pull request template](./.github/PULL_REQUEST_TEMPLATE.md). It is loaded automatically when you open a PR: fill in every section and do not remove any of them.

> [!WARNING]
> A pull request that does not follow the template may be refused (closed without review).

| Section | Content |
|---------|---------|
| **Description** | Short summary of the changes and why they are needed. Link the issue with `Closes #<number>` |
| **Type of change** | Bug fix, new feature, breaking change, refactoring, documentation, infrastructure / CI-CD or performance |
| **Changes made** | List of the main modifications |
| **How to test** | How a reviewer can verify the implementation |
| **Self-check** | Checklist to complete before requesting a review (style, self-review, tests, documentation, no secrets in the diff, no new linter warnings, PR title format) |

**Guidelines**
- One PR per feature or fix, keep the scope focused
- The PR title follows the [Commits](#commits) convention (`type(scope): short description`), since the PR is squash-merged
- All commits are signed (see [Commits](#commits))
- Update documentation if the PR affects the public API
- All checks must pass before requesting a review (`make ci`, and `make test` for the tests alone)

---

## Issues

Issues must be opened from one of the issue templates. Blank issues are disabled.

| Template | Use it to | Title prefix |
|----------|-----------|--------------|
| `bug_report` | Report a bug | `[Bug]` |
| `feature_request` | Suggest a feature or an improvement | `[Feature]` |
| `documentation` | Report an error or suggest an improvement in the documentation | `[Docs]` |

Before opening an issue, check the [documentation](./docs/README.md) and the existing issues: each template asks you to confirm that your issue is not a duplicate. For questions and discussions, join the [Discord server](https://discord.gg/6yNACMQqV9).

### Reporting Bugs

Use the **Bug Report** template. Required fields:
- Description
- Steps to reproduce
- Expected behavior
- Actual behavior (with the exact error message if any)
- Affected version of the project
- Operating system

Logs and screenshots are optional but help a lot.

### Suggesting Features

Use the **Feature Request** template. It asks for:
- The problem to solve
- The solution you propose
- The alternatives you considered (optional)
- The impact for you: nice to have, regular use, or blocking your workflow

### Documentation Issues

Use the **Documentation** template to report a typo, inaccurate information, missing content, an outdated example, confusing wording or a broken link. Indicate the affected page and, if you can, the corrected text.

### Security Vulnerabilities

Do **not** open a public issue for a security vulnerability. Follow the [security policy](./.github/SECURITY.md) instead.

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

---

## See Also

- [Architecture](./docs/contributing/architecture.md)
- [Testing](./docs/contributing/testing.md)
- [Documentation index](./docs/README.md)
- [Pull request template](./.github/PULL_REQUEST_TEMPLATE.md)
- [Security policy](./.github/SECURITY.md)
- [Code of Conduct](./.github/CODE_OF_CONDUCT.md)
- [Effective Go](https://go.dev/doc/effective_go)
- [Karma Commit Convention](http://karma-runner.github.io/6.4/dev/git-commit-msg.html)