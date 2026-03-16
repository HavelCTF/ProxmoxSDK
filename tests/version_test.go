package tests

import (
	"net/http"
	"testing"

	"github.com/HavelCTF/ProxmoxSDK/pkg/proxmox"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

const (
	TestURI string = "http://test.localhost"
)

func version(baseUrl string) {
	versionJSON := `
{
    "data": {
        "repoid": "9a1b2c3d",
        "release": "9.1",
        "version": "9.1-1"
    }
}`
	gock.New(baseUrl).
		Get("/version").
		Reply(200).
		JSON(versionJSON)
}

func TestGetVersion(t *testing.T) {
	defer gock.Off()
	httpClient := http.DefaultClient
	gock.InterceptClient(httpClient)

	mockClient := proxmox.NewClient(TestURI, "testoken", "testuuid",
		proxmox.WithHTTPClient(httpClient))
	version(mockClient.BaseURL())
	v, err := mockClient.GetVersion()

	assert.Nil(t, err)
	assert.Equal(t, "9.1-1", v.Data.Version)
}
