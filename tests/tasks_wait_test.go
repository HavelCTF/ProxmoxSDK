package tests

import (
	"context"
	"testing"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/tests/mocks"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestTaskWaitTransitions(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.TaskWaitSequence(mockClient.BaseURL())

	upid := "UPID:node1:00000099:00000099:00000099:test:waitme:root@pam:"
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := mockClient.Node("node1").Tasks(upid).Wait(ctx, 20*time.Millisecond)

	assert.NoError(t, err)
	assert.Equal(t, "stopped", resp.Data.Status)
	assert.Equal(t, "OK", resp.Data.ExitStatus)
	assert.Equal(t, upid, resp.Data.UPID)
	assert.True(t, gock.IsDone())
}

func TestTaskWaitContextCancel(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.TaskWaitRunningPersist(mockClient.BaseURL())

	upid := "UPID:node1:00000098:00000098:00000098:test:cancelme:root@pam:"
	ctx, cancel := context.WithTimeout(context.Background(), 80*time.Millisecond)
	defer cancel()

	_, err := mockClient.Node("node1").Tasks(upid).Wait(ctx, 30*time.Millisecond)

	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestTaskWaitDefaultPollInterval(t *testing.T) {
	defer gock.Off()
	mockClient := initTestClient()
	mocks.TaskWaitSequence(mockClient.BaseURL())

	upid := "UPID:node1:00000099:00000099:00000099:test:waitme:root@pam:"
	// Use a generous timeout; 0 pollInterval should fall back to the 1s default,
	// and the second mock returns "stopped" immediately on the second hit.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := mockClient.Node("node1").Tasks(upid).Wait(ctx, 0)

	assert.NoError(t, err)
	assert.Equal(t, "stopped", resp.Data.Status)
	assert.True(t, gock.IsDone())
}
