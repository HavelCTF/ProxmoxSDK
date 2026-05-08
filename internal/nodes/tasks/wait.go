package tasks

import (
	"context"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/types"
)

// defaultWaitPollInterval is used when Wait is called with a zero pollInterval.
const defaultWaitPollInterval = 1 * time.Second

// Wait polls GetTaskStatus on the configured interval until the Proxmox task
// transitions out of the "running" state, then returns the final status.
//
// A zero pollInterval falls back to defaultWaitPollInterval (1s). Callers
// should attach an absolute deadline via ctx (e.g. context.WithTimeout) — Wait
// returns ctx.Err() as soon as the context is done.
func (s *TaskService) Wait(ctx context.Context, pollInterval time.Duration) (*types.NodeTaskStatusResponse, error) {
	if pollInterval <= 0 {
		pollInterval = defaultWaitPollInterval
	}

	// First check is immediate so already-finished tasks return without delay.
	for {
		status, err := s.GetTaskStatus()
		if err != nil {
			return nil, err
		}
		if status.Data.Status != "running" {
			return status, nil
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(pollInterval):
		}
	}
}
