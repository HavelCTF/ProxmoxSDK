package types

type NodeStatus string

const (
	Unknown NodeStatus = "unknown"
	Online  NodeStatus = "online"
	Offline NodeStatus = "offline"
)

// NodesResponse represents the response from the Proxmox /nodes endpoint.
type NodesResponse struct {
	Data []NodesData `json:"data"`
}

// Node contains Proxmox node informations.
type NodesData struct {
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

// NodeResponse represents the response from the Proxmox /nodes/{node}
// endpoint.
type NodeResponse struct {
	Data []NodeData `json:"data"`
}

// NodeProperty contains Proxmox node property name.
type NodeData struct {
	Name string `json:"name"`
}

type NodeTasksResponse struct {
	Data []NodeTasksData `json:"data"`
}

type NodeTaskBase struct {
	ID        string `json:"id"`
	Node      string `json:"node"`
	PID       int    `json:"pid"`
	PStart    int    `json:"pstart"`
	StartTime int    `json:"starttime"`
	Type      string `json:"type"`
	UPID      string `json:"upid"`
	User      string `json:"user"`
}

type NodeTasksData struct {
	NodeTaskBase
	EndTime int    `json:"endtime,omitempty"`
	Status  string `json:"status,omitempty"`
}

type NodeTaskStatusResponse struct {
	Data NodeTaskStatusData `json:"data"`
}

type NodeTaskStatusData struct {
	NodeTaskBase
	Status     string `json:"status"`
	ExitStatus string `json:"exitstatus,omitempty"`
}

type NodeTaskDeleteResponse struct {
	Data string `json:"data"`
}
