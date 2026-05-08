package tests

import (
	"testing"

	"github.com/HavelCTF/ProxmoxSDK/tests/mocks"
	"github.com/HavelCTF/ProxmoxSDK/types"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestLXCGet(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.LXCStatusCurrent(mockClient.BaseURL())

	resp, err := mockClient.Node("node1").LXC(100).Get()

	assert.NoError(t, err)
	assert.Equal(t, types.Running, resp.Data.Status)
	assert.Equal(t, "ct-test-1", resp.Data.Name)
	assert.Equal(t, 100, resp.Data.VMID)
	assert.Equal(t, 12345, resp.Data.Uptime)
	assert.True(t, gock.IsDone())
}

func TestLXCGetStopped(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.LXCStatusCurrentStopped(mockClient.BaseURL())

	resp, err := mockClient.Node("node1").LXC(101).Get()

	assert.NoError(t, err)
	assert.Equal(t, types.Stopped, resp.Data.Status)
	assert.Equal(t, 0, resp.Data.Uptime)
	assert.True(t, gock.IsDone())
}
