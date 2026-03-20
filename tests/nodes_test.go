package tests

import (
	"fmt"
	"testing"

	"github.com/HavelCTF/ProxmoxSDK/pkg/proxmox"
	"github.com/HavelCTF/ProxmoxSDK/tests/mocks"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestNodeService(t *testing.T) {
	client := proxmox.NewClient(testURI, testToken, testUUID)
	nodeService := client.Node("testnode")

	assert.Equal(t, "testnode", nodeService.Node())
}

func TestGetNodes(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.Nodes(mockClient.BaseURL())

	nodes, err := mockClient.GetNodes()
	assert.NoError(t, err)

	for idx, node := range nodes.Data {
		assert.Equal(t, fmt.Sprintf("node%d", idx+1), node.Node)
	}
	assert.True(t, gock.IsDone())
}
