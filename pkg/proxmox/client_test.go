package proxmox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient("https://pve.example.com:8006", "root@pam!test=secret", "test-uuid")

	assert.Equal(t, "https://pve.example.com:8006/api2/json", c.BaseURL())
	assert.Equal(t, "test-uuid", c.UUID())
}
