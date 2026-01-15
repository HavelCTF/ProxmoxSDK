package proxmox_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/HavelCTF/ProxmoxSDK/go/proxmox"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Generic spec structures
type ClientTestCase struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Input       struct {
		BaseURL  string `json:"baseURL"`
		APIToken string `json:"apiToken"`
		UUID     string `json:"uuid"`
	} `json:"input"`
	Expected struct {
		BaseURL string `json:"baseURL"`
		UUID    string `json:"uuid"`
	} `json:"expected"`
}

type MockRetry struct {
	Attempt    int                    `json:"attempt"`
	StatusCode int                    `json:"statusCode"`
	Body       map[string]interface{} `json:"body"`
}

type MockConfig struct {
	StatusCode int                    `json:"statusCode"`
	Body       map[string]interface{} `json:"body"`
	Retries    []MockRetry            `json:"retries"`
}

type VersionTestCase struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Input       struct {
		BaseURL  string `json:"baseURL"`
		APIToken string `json:"apiToken"`
		UUID     string `json:"uuid"`
	} `json:"input"`
	Mock             MockConfig             `json:"mock"`
	Expected         map[string]interface{} `json:"expected"`
	ExpectError      bool                   `json:"expectError"`
	ErrorContains    string                 `json:"errorContains"`
	ExpectedAttempts int                    `json:"expectedAttempts"`
}

func loadSpec(t *testing.T, filename string) []byte {
	path := filepath.Join("..", "test-specs", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read spec %s: %v", filename, err)
	}
	return data
}

func loadAllSpecs(t *testing.T) map[string][]byte {
	specsDir := filepath.Join("..", "test-specs")
	specs := make(map[string][]byte)

	entries, err := os.ReadDir(specsDir)
	if err != nil {
		t.Fatalf("failed to read test-specs directory: %v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		specName := strings.TrimSuffix(entry.Name(), ".json")
		specs[specName] = loadSpec(t, entry.Name())
	}

	return specs
}

func TestAllSpecs(t *testing.T) {
	allSpecs := loadAllSpecs(t)

	// Run client specs
	if data, exists := allSpecs["client"]; exists {
		t.Run("ClientSpecs", func(t *testing.T) {
			var cases []ClientTestCase
			if err := json.Unmarshal(data, &cases); err != nil {
				t.Fatalf("failed to parse client spec: %v", err)
			}

			for _, tc := range cases {
				t.Run(tc.Name, func(t *testing.T) {
					client := proxmox.NewClient(
						tc.Input.BaseURL,
						tc.Input.APIToken,
						tc.Input.UUID,
					)

					require.NotNil(t, client)
					assert.Equal(t, tc.Expected.UUID, client.GetUUID())
					assert.Equal(t, tc.Expected.BaseURL, client.GetBaseURL())
				})
			}
		})
	}

	// Run version specs
	if data, exists := allSpecs["version"]; exists {
		t.Run("VersionSpecs", func(t *testing.T) {
			var cases []VersionTestCase
			if err := json.Unmarshal(data, &cases); err != nil {
				t.Fatalf("failed to parse version spec: %v", err)
			}

			for _, tc := range cases {
				t.Run(tc.Name, func(t *testing.T) {
					var callCount int

					// Create mock server
					server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						callCount++

						// Verify request
						assert.Equal(t, tc.Input.APIToken, r.Header.Get("Authorization"))
						assert.Contains(t, r.URL.Path, "/api2/json/version")

						w.Header().Set("Content-Type", "application/json")

						// Handle retry scenarios
						if len(tc.Mock.Retries) > 0 {
							for _, retry := range tc.Mock.Retries {
								if retry.Attempt == callCount {
									w.WriteHeader(retry.StatusCode)
									if retry.Body != nil {
										json.NewEncoder(w).Encode(retry.Body)
									}
									return
								}
							}
							w.WriteHeader(http.StatusInternalServerError)
							return
						}

						// Single response
						w.WriteHeader(tc.Mock.StatusCode)
						if tc.Mock.Body != nil {
							json.NewEncoder(w).Encode(tc.Mock.Body)
						}
					}))
					defer server.Close()

					// Create client
					client := proxmox.NewClient(server.URL, tc.Input.APIToken, tc.Input.UUID)

					// Execute version request
					result, err := client.Version()

					// Verify expectations
					if tc.ExpectError {
						require.Error(t, err)
						if tc.ErrorContains != "" {
							assert.Contains(t, err.Error(), tc.ErrorContains)
						}
					} else {
						require.NoError(t, err)
						require.NotNil(t, result)

						if tc.Expected != nil {
							assert.Equal(t, tc.Expected["version"], result.Data.Version)
							assert.Equal(t, tc.Expected["release"], result.Data.Release)
							assert.Equal(t, tc.Expected["repoid"], result.Data.RepoID)
						}
					}

					// Verify attempt count
					if tc.ExpectedAttempts > 0 {
						assert.Equal(t, tc.ExpectedAttempts, callCount)
					}
				})
			}
		})
	}
}
