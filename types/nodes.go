package types

type NodeStatus string

const (
	Unknown NodeStatus = "unknown"
	Online  NodeStatus = "online"
	Offline NodeStatus = "offline"
)

type NodesResponse struct {
	Nodes []Node `json:"data"`
}

type Node struct {
	Node           string     `json:"node"`
	Status         NodeStatus `json:"status"`
	Uptime         int        `json:"uptime,omitempty"`
	SslFingerprint string     `json:"ssl_fingerprint,omitempty"`
	CPU            float32    `json:"cpu,omitempty"`
	Level          string     `json:"level,omitempty"`
	MaxCPU         int        `json:"maxcpu,omitempty"`
	MaxMEM         int        `json:"maxmem,omitempty"`
	MEM            int        `json:"mem,omitempty"`
}
