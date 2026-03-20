package tests

import (
	"testing"

	"github.com/HavelCTF/ProxmoxSDK/pkg/proxmox"
	"github.com/HavelCTF/ProxmoxSDK/tests/mocks"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestTaskService(t *testing.T) {
	client := proxmox.NewClient(testURI, testToken, testUUID)
	taskService := client.Node("node1").Tasks("UPID:node1:00000001:00000001:00000001:test:running:root@pam:")

	assert.Equal(t, "node1", taskService.Node())
	assert.Equal(t, "UPID:node1:00000001:00000001:00000001:test:running:root@pam:", taskService.UPID())
}

func TestGetTasks(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.Tasks(mockClient.BaseURL())

	tasks, err := mockClient.Node("node1").GetTasks()

	assert.NoError(t, err)
	for _, task := range tasks.Data {
		assert.Equal(t, "node1", task.Node)
		assert.Equal(t, "aptupdate", task.Type)
	}
	assert.True(t, gock.IsDone())
}

func TestGetTasksStatus(t *testing.T) {
	defer gock.Off()
	testTable := []struct {
		name         string
		upid         string
		statusResult string
	}{
		{"task running", "UPID:node1:00000001:00000001:00000001:test:running:root@pam:", "running"},
		{"task completed", "UPID:node1:00000002:00000002:00000002:test:completed:root@pam:", "stopped"},
		{"task failed", "UPID:node1:00000003:00000003:00000003:test:failed:root@pam:", "stopped"},
	}
	mockClient := initTestClient()
	mocks.TasksStatus(mockClient.BaseURL())
	for _, test := range testTable {
		resp, err := mockClient.Node("node1").Tasks(test.upid).GetTaskStatus()

		assert.NoError(t, err)
		assert.Equal(t, test.statusResult, resp.Data.Status)
		assert.Equal(t, test.upid, resp.Data.UPID)
	}
}

func TestDeleteTask(t *testing.T) {
	defer gock.Off()
	testTable := []struct {
		name         string
		upid         string
		expectError  bool
		errorMessage string
	}{
		{"stop task", "UPID:node1:00000001:00000001:00000001:test:running:root@pam:", false, ""},
		{"task does not exist", "UPID:node1:00000001:00000001:00000003:test::root@pam:", true, "no such task"},
		{"upid format error", "UPID:node1:00000001:0000001:0000000Z:test::root@pam:", true, "unable to parse worker upid"},
	}
	defer gock.Off()
	mockClient := initTestClient()
	mocks.StopTask(mockClient.BaseURL())

	for _, test := range testTable {
		resp, err := mockClient.Node("node1").Tasks(test.upid).DeleteTask()

		if test.expectError == true {
			assert.ErrorContains(t, err, test.errorMessage)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, "", resp.Data)
		}
	}
}
