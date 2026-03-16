package proxmox

import "testing"

func TestNewClient(t *testing.T) {
	c := NewClient("https://pve.example.com:8006", "root@pam!test=secret", "test-uuid")

	if c.BaseURL() != "https://pve.example.com:8006/api2/json" {
		t.Errorf("expected base URL %q, got %q", "https://pve.example.com:8006/api2/json", c.BaseURL())
	}
	if c.UUID() != "test-uuid" {
		t.Errorf("expected UUID %q, got %q", "test-uuid", c.UUID())
	}
}
