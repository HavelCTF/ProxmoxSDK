# Test Specifications

This directory contains language-agnostic test specifications that define expected behavior for the ProxmoxSDK implementations in both TypeScript and Go.

## Overview

Test specifications are written in YAML format and define:
- Test scenarios with inputs and expected outputs
- Mock server responses
- Error conditions
- Retry behavior
- Timeout handling

Both the TypeScript and Go test suites consume these specifications to ensure identical behavior across implementations.

## Format

Each test specification file contains:

```yaml
version: "1.0"
tests:
  - name: "Test name"
    description: "What this test verifies"
    scenario:
      input:
        baseURL: "https://proxmox.example.com"
        apiToken: "test-token"
        uuid: "test-uuid"
      mock:
        endpoint: "/api2/json/version"
        method: "GET"
        response:
          status: 200
          body:
            version: "8.0.3"
            release: "8.0"
            repoid: "abc123"
      expected:
        success: true
        output:
          version: "8.0.3"
          release: "8.0"
          repoid: "abc123"
```

## Running Tests

### TypeScript
```bash
cd typescript
npm test
```

### Go
```bash
cd go
go test ./...
```

Both test runners will automatically load and execute the shared specifications.

## Test Categories

- `client-spec.yaml` - Client initialization tests
- `version-spec.yaml` - Version endpoint tests  
- `error-handling-spec.yaml` - Error handling and retry logic tests
