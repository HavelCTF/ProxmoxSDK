package tests

import (
	"testing"

	"github.com/HavelCTF/ProxmoxSDK/tests/mocks"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestClusterGetNextId(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.ClusterNextId(mockClient.BaseURL())

	nextId, err := mockClient.Cluster().GetNextId()

	assert.NoError(t, err)
	assert.Equal(t, "100", nextId.VMID)
	assert.True(t, gock.IsDone())
}

func TestGetTasks(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.ClusterTasks(mockClient.BaseURL())

	tasks, err := mockClient.Cluster().GetTasks()
	assert.NoError(t, err)
	assert.Equal(t, "UPID:node1:00000001:00000001:00000001:test:running:root@pam:", tasks.Data[0].UPID)
	assert.True(t, gock.IsDone())
}
