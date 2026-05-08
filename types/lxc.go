package types

import (
	"fmt"
	"net/url"
)

type LXCStatus string

const (
	Stopped LXCStatus = "stopped"
	Running LXCStatus = "running"
)

// LXCsResponse maps to the GET /nodes/{node}/lxc API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc
type LXCsResponse struct {
	LXCs []LXCsData `json:"data"`
}

type LXCsData struct {
	Status             LXCStatus `json:"status"`
	VMID               int       `json:"vmid"`
	CPU                float32   `json:"cpu,omitempty"`
	CPUS               float32   `json:"cpus,omitempty"`
	Disk               int       `json:"disk,omitempty"`
	DiskRead           int       `json:"diskread,omitempty"`
	DiskWrite          int       `json:"diskwrite,omitempty"`
	Lock               string    `json:"lock,omitempty"`
	MaxDisk            int       `json:"maxdisk,omitempty"`
	MaxMem             int       `json:"maxmem,omitempty"`
	MaxSwap            int       `json:"maxswap,omitempty"`
	Mem                int       `json:"mem,omitempty"`
	Name               string    `json:"name,omitempty"`
	NetIn              int       `json:"netin,omitempty"`
	NetOut             int       `json:"netout,omitempty"`
	PressureCPUSome    string    `json:"pressurecpusome,omitempty"`
	PressureIOFull     string    `json:"pressureiofull,omitempty"`
	PressureIOSome     string    `json:"pressureiosome,omitempty"`
	PressureMemoryFull string    `json:"pressurememoryfull,omitempty"`
	PressureMemorySome string    `json:"pressurememorysome,omitempty"`
	Tags               string    `json:"tags,omitempty"`
	Template           bool      `json:"template,omitempty"`
	Uptime             int       `json:"uptime,omitempty"`
}

// CreateLXCData maps to the POST /nodes/{node}/lxc API request content
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc
type CreateLXCData struct {
	OSTemplate string       `url:"ostemplate"`
	VMID       int          `url:"vmid"`
	Features   *LXCFeatures `url:"features,omitempty"`
}

type LXCFeatures struct {
	Nesting bool
}

// EncodeValues implements query.Encoder.
func (f LXCFeatures) EncodeValues(key string, v *url.Values) error {
	v.Set(key, f.Encode())
	return nil
}

func (f LXCFeatures) Encode() string {
	val := 0
	if f.Nesting {
		val = 1
	}
	return fmt.Sprintf("nesting=%d", val)
}

// CloneLXCData maps to the POST /nodes/{node}/lxc/{vmid}/clone API request content
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/clone
type CloneLXCData struct {
	NewId  int     `url:"newid"`
	Target *string `url:"target,omitempty"`
}

// CreateLXCResponse maps to the POST /nodes/{node}/lxc API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc
type CreateLXCResponse TaskBaseResponse

// CloneLXCResponse maps to the POST /nodes/{node}/lxc/{vmid}/clone API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/clone
type CloneLXCResponse TaskBaseResponse

// DeleteLXCResponse maps to the DELETE /nodes/{node}/lxc/{vmid} API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}
type DeleteLXCResponse TaskBaseResponse

// StartLXCResponse maps to the POST /nodes/{node}/lxc/{vmid}/status/start API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/status/start
type StartLXCResponse TaskBaseResponse

// StopLXCResponse maps to the POST /nodes/{node}/lxc/{vmid}/status/stop API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/status/stop
type StopLXCResponse TaskBaseResponse

// LXCStatusResponse maps to the GET /nodes/{node}/lxc/{vmid}/status/current API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/status/current
type LXCStatusResponse struct {
	Data LXCStatusData `json:"data"`
}

type LXCStatusData struct {
	Name      string    `json:"name,omitempty"`
	Status    LXCStatus `json:"status"`
	VMID      int       `json:"vmid,omitempty"`
	Uptime    int       `json:"uptime,omitempty"`
	CPUs      float32   `json:"cpus,omitempty"`
	CPU       float32   `json:"cpu,omitempty"`
	Mem       int       `json:"mem,omitempty"`
	MaxMem    int       `json:"maxmem,omitempty"`
	Disk      int       `json:"disk,omitempty"`
	MaxDisk   int       `json:"maxdisk,omitempty"`
	Swap      int       `json:"swap,omitempty"`
	MaxSwap   int       `json:"maxswap,omitempty"`
	NetIn     int       `json:"netin,omitempty"`
	NetOut    int       `json:"netout,omitempty"`
	DiskRead  int       `json:"diskread,omitempty"`
	DiskWrite int       `json:"diskwrite,omitempty"`
	HA        any       `json:"ha,omitempty"`
	Lock      string    `json:"lock,omitempty"`
	Tags      string    `json:"tags,omitempty"`
	Type      string    `json:"type,omitempty"`
}
