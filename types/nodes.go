package types

type NodeStatus string

const (
	Unknown NodeStatus = "unknown"
	Online  NodeStatus = "online"
	Offline NodeStatus = "offline"
)

// TaskBaseResponse contains the UPID returned by Proxmox for asynchronous operations
// like container/VM creation, deletion, or modification.
type TaskBaseResponse struct {
	UPID string `json:"data"`
}

// NodesResponse maps to the GET /nodes API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes
type NodesResponse struct {
	Data []NodesData `json:"data"`
}

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

// NodeResponse maps to the GET /nodes/{node} API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}
type NodeResponse struct {
	Data []NodeData `json:"data"`
}

type NodeData struct {
	Name string `json:"name"`
}

// NodeTasksResponse maps to the GET /nodes/{node}/tasks API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/tasks
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

// NodeTaskStatusResponse maps to the GET /nodes/{node}/tasks/{upid}/status API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/tasks/{upid}/status
type NodeTaskStatusResponse struct {
	Data NodeTaskStatusData `json:"data"`
}

type NodeTaskStatusData struct {
	NodeTaskBase
	Status     string `json:"status"`
	ExitStatus string `json:"exitstatus,omitempty"`
}

// NodeTaskDeleteResponse maps to the DELETE /nodes/{node}/tasks/{upid} API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/tasks/{upid}
type NodeTaskDeleteResponse struct {
	Data string `json:"data"`
}
