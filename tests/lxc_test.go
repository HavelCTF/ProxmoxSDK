package tests

import (
	"fmt"
	"testing"

	"github.com/HavelCTF/ProxmoxSDK/pkg/proxmox"
	"github.com/HavelCTF/ProxmoxSDK/tests/mocks"
	"github.com/HavelCTF/ProxmoxSDK/types"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestLXCService(t *testing.T) {
	client := proxmox.NewClient(testURI, testToken, testUUID)
	lxcService := client.Node("node1").LXC(101)

	assert.Equal(t, "node1", lxcService.Node())
	assert.Equal(t, 101, lxcService.VMID())
}

func TestGetLXCs(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.LXC(mockClient.BaseURL())

	lxcs, err := mockClient.Node("node1").GetLXCs()

	assert.NoError(t, err)
	for idx, lxc := range lxcs.LXCs {
		assert.Equal(t, 100+idx, lxc.VMID)
		assert.Equal(t, fmt.Sprintf("ct-test-%d", idx+1), lxc.Name)
	}
	assert.True(t, gock.IsDone())
}

func TestPostLXC(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.PostLXC(mockClient.BaseURL())

	result, err := mockClient.Node("node1").PostLXC(types.CreateLXCData{
		OSTemplate: "path-to-template",
		VMID:       101,
		Features: &types.LXCFeatures{
			Nesting: true,
		},
	})

	assert.NoError(t, err)
	assert.Equal(t, "UPID:node1:00001234:00005678:5A3B7C8D:vzcreate:101:root@pam:", result.UPID)
	assert.True(t, gock.IsDone())
}

func TestLXCStart(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.LXCStart(mockClient.BaseURL())

	result, err := mockClient.Node("node1").LXC(101).Status().StartLXC()

	assert.NoError(t, err)
	assert.Equal(t, "UPID:node1:00001234:00005678:5A3B7C8D:vzstart:101:root@pam:", result.UPID)
	assert.True(t, gock.IsDone())
}

func TestLXCStop(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.LXCStop(mockClient.BaseURL())

	result, err := mockClient.Node("node1").LXC(101).Status().StopLXC()

	assert.NoError(t, err)
	assert.Equal(t, "UPID:node1:00001234:00005678:5A3B7C8D:vzstop:101:root@pam:", result.UPID)
	assert.True(t, gock.IsDone())
}

func TestDeleteLXC(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.DeleteLXC(mockClient.BaseURL())

	result, err := mockClient.Node("node1").LXC(101).DeleteLXC()

	assert.NoError(t, err)
	assert.Equal(t, "UPID:node1:00001234:00005678:5A3B7C8D:vzdestroy:101:root@pam:", result.UPID)
	assert.True(t, gock.IsDone())
}

func TestCloneLXC(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.CloneLXC(mockClient.BaseURL())

	target := "node1"
	result, err := mockClient.Node("node1").LXC(101).CloneLXC(types.CloneLXCData{
		NewId:  102,
		Target: &target,
	})

	assert.NoError(t, err)
	assert.Equal(t, "UPID:node1:00001234:00005678:5A3B7C8D:vzmigrate:101:root@pam:", result.UPID)
	assert.True(t, gock.IsDone())
}
