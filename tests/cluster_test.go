// SPDX-License-Identifier: MIT
// Copyright (c) 2026 HavelCTF

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
	err := mocks.InitMocker(mockClient.BaseURL(), "GET", "/cluster/nextid", 200, "cluster/nextid.json")
	if err != nil {
		assert.NoError(t, err, "Failed to initialize mocker: %v", err)
	}

	nextId, err := mockClient.Cluster().GetNextId()
	assert.NoError(t, err)
	assert.Equal(t, "100", nextId.VMID)
	assert.True(t, gock.IsDone())
}

func TestClusterGetTasks(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	err := mocks.InitMocker(mockClient.BaseURL(), "GET", "/cluster/tasks", 200, "cluster/tasks.json")
	if err != nil {
		assert.NoError(t, err, "Failed to initialize mocker: %v", err)
	}

	tasks, err := mockClient.Cluster().GetTasks()
	assert.NoError(t, err)
	assert.Equal(t, "UPID:node1:00000001:00000001:00000001:test:running:root@pam:", tasks.Data[0].UPID)
	assert.True(t, gock.IsDone())
}
