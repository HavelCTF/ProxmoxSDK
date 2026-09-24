# Testing

> ProxmoxSDK uses black-box testing to validate all operations exposed to the client.

## Table of Contents

- [Philosophy](#philosophy)
- [Structure](#structure)
- [Running Tests](#running-tests)
- [Writing Tests](#writing-tests)
  - [Test Functions](#test-functions)
  - [Table-Driven Tests](#table-driven-tests)
  - [Mocks](#mocks)
- [Coverage](#coverage)
- [See Also](#see-also)

---

## Philosophy

Tests target only the public API of the SDK: what the client can call. Internal implementation details are not tested directly. This is known as black-box testing: the test file lives in a separate `_test` package and can only access exported identifiers. 

The goal is to verify that each service method returns the expected response given a controlled HTTP mock, without requiring a real Proxmox instance.

---

## Structure
```
tests/
├── config.go           # shared client initialization
├── mocks/              # gock HTTP mocks per service
│   ├── cluster.go
│   ├── ...
├── cluster_test.go
├── ...
```

Each service has one `{service}_test.go` file and one corresponding mock file in `mocks/`.

**`config.go`** provides a shared test client and constants used across all test files:
```go
const (
    testURI   string = "http://test.localhost"
    testToken string = "PVEAPIToken=root@pam!testtoken=testtoken"
    testUUID  string = "test-0"
)

func initTestClient() *proxmox.Client {
    httpClient := http.DefaultClient
    gock.InterceptClient(httpClient)

    return proxmox.NewClient(testURI, testToken, testUUID,
        proxmox.WithHTTPClient(httpClient))
}
```

---

## Running Tests
```bash
# Run all tests
make test

# Run with coverage report
make test-coverage

# Run manually with verbose output
go test -v ./...
```

---

## Writing Tests

### Test Functions

Test functions follow the `TestXxx` naming convention where `Xxx` starts with a capital letter and identifies what is being tested.  Name tests after the method and the scenario being verified:
```go
func TestGetVersion(t *testing.T) { ... }
func TestGetNextId(t *testing.T) { ... }
func TestGetTasks(t *testing.T) { ... }
```

Each test function initializes a fresh client and defers `gock.Off()` to clean up interceptors:
```go
func TestGetVersion(t *testing.T) {
    client := initTestClient()
    defer gock.Off()

    mocks.Version(testURI)

    version, err := client.GetVersion()
    assert.NoError(t, err)
    assert.Equal(t, "8.1.4", version.Data.Version)
}
```

### Table-Driven Tests

Use table-driven tests when a method needs to be verified across multiple scenarios. Each table entry is a complete test case with inputs and expected results. 
```go
func TestGetVersion(t *testing.T) {
	testTable := []struct {
		name string
		mock func(string)
	}{
		{"nominal", mocks.Version},
		{"retry on 500", mocks.VersionWithServiceUnavailable},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			defer gock.Off()
			mockClient := initTestClient()
			test.mock(mockClient.BaseURL())

			v, err := mockClient.GetVersion()

			assert.NoError(t, err)
			assert.Equal(t, "9.1-1", v.Data.Version)
			assert.Equal(t, "9a1b2c3d", v.Data.RepoID)
			assert.Equal(t, "9.1", v.Data.Release)
			assert.True(t, gock.IsDone())
		})
	}
}
```

### Mocks

Mocks are defined in `tests/mocks/` using [gock](https://github.com/h2non/gock). Each mock file exposes one function per scenario:
```go
// mocks/version.go

// nominal case
func Version(baseUrl string) {
    gock.New(baseUrl).Get("/version").Reply(200).JSON(versionJSON)
}

// retry scenario: two failures before success
func VersionWithServiceUnavailable(baseUrl string) {
    gock.New(baseUrl).Get("/version").Reply(500)
    gock.New(baseUrl).Get("/version").Reply(500)
    gock.New(baseUrl).Get("/version").Reply(200).JSON(versionJSON)
}
```

**Rules:**
- One mock function per scenario (`Version`, `VersionWithServiceUnavailable`, ...)
- Multiple `gock` interceptors can be chained in a single function to simulate retry behavior
- Fixtures defined in the same file as their mock functions
- Mock functions take `baseUrl string` as their only parameter

---

## Coverage

No minimum coverage threshold is enforced. The goal is the highest coverage achievable through black-box testing. Every exported method should have at least one test covering its nominal case.

To visualize coverage:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

---

## See Also

- [Architecture](./architecture.md)
- [testify documentation](https://github.com/stretchr/testify)
- [gock documentation](https://github.com/h2non/gock)
- [Go testing package](https://pkg.go.dev/testing)
- [Effective Go](https://go.dev/doc/effective_go)