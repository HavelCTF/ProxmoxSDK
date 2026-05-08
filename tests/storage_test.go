package tests

import (
	"strings"
	"testing"

	"github.com/HavelCTF/ProxmoxSDK/pkg/proxmox"
	"github.com/HavelCTF/ProxmoxSDK/tests/mocks"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestStorageService(t *testing.T) {
	client := proxmox.NewClient(testURI, testToken, testUUID)
	storageSvc := client.Node("node1").Storage("local")

	assert.Equal(t, "node1", storageSvc.Node())
	assert.Equal(t, "local", storageSvc.Storage())
}

func TestStorageGetContent(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.StorageContent(mockClient.BaseURL())

	resp, err := mockClient.Node("node1").Storage("local").GetContent("vztmpl")

	assert.NoError(t, err)
	assert.Len(t, resp.Data, 2)
	assert.Equal(t, "local:vztmpl/debian-12-standard_12.7-1_amd64.tar.zst", resp.Data[0].VolID)
	assert.Equal(t, "vztmpl", resp.Data[0].Content)
	assert.Equal(t, "tzst", resp.Data[0].Format)
	assert.True(t, gock.IsDone())
}

func TestStorageUploadTemplate(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.StorageUpload(mockClient.BaseURL())

	body := strings.NewReader("not-a-real-template-payload")
	resp, err := mockClient.Node("node1").Storage("local").UploadTemplate(
		"debian-12-standard_12.7-1_amd64.tar.zst",
		body,
		int64(body.Len()),
	)

	assert.NoError(t, err)
	assert.Equal(t, "UPID:node1:00001234:00005678:5A3B7C8D:imgcopy::root@pam:", resp.UPID)
	assert.True(t, gock.IsDone())
}
