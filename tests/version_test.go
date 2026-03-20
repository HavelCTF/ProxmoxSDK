package tests

import (
	"testing"

	"github.com/HavelCTF/ProxmoxSDK/tests/mocks"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

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
