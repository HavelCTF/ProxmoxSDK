package types

type LXCStatus string

const (
	Stopped LXCStatus = "stopped"
	Running LXCStatus = "running"
)

// type LXCArch string

// const (
// 	Amd64   LXCArch = "amd64"
// 	I386    LXCArch = "i386"
// 	Arm64   LXCArch = "arm64"
// 	Armhf   LXCArch = "armhf"
// 	Riscv32 LXCArch = "riscv32"
// 	Riscv64 LXCArch = "riscv64"
// )

// type LXCConsoleMode string

// const (
// 	Shell   LXCConsoleMode = "shell"
// 	Console LXCConsoleMode = "console"
// 	TTY     LXCConsoleMode = "tty"
// )

// type LXCDeviceMode uint32

// const (
// 	OwnerR        LXCDeviceMode = 0400 // r--------
// 	OwnerRW       LXCDeviceMode = 0600 // rw-------
// 	OwnerRWX      LXCDeviceMode = 0700 // rwx------
// 	OwnerGroupR   LXCDeviceMode = 0440 // r--r-----
// 	OwnerGroupRW  LXCDeviceMode = 0660 // rw-rw----
// 	OwnerGroupRWX LXCDeviceMode = 0770 // rwxrwx---
// 	AllR          LXCDeviceMode = 0444 // r--r--r--
// 	AllRW         LXCDeviceMode = 0666 // rw-rw-rw-
// 	AllRWX        LXCDeviceMode = 0777 // rwxrwxrwx
// )

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
	Node       string `url:"node"`
	OSTemplate string `url:"ostemplate"`
	VMID       int    `url:"vmid"`
	// Arch           *LXCArch
	// BandwidthLimit *float32
	// ConsoleMode    *LXCConsoleMode
	// Console        *bool
	// Cores          *int
	// CPULimit       *float32
	// CPUUnits       *int
	// Debug          *bool
	// Description    *string
	// Device         *[]LXCDevice
	// EntryPoint     *string
	// Env            *map[string]string
	// Features       *LXCFeatures
}

// type LXCDevice struct {
// 	Path      string
// 	DenyWrite *bool
// 	GID       *int
// 	Mode      *LXCDeviceMode
// 	UID       *int
// }

// type LXCFeatures struct {
// 	ForceRWSys bool
// 	Fuse       *bool
// 	KeyCTL     *bool
// 	MkNod      *bool
// 	Mount      *[]string
// 	Nesting    *bool
// }

// CreateLXCResponse maps to the POST /nodes/{node}/lxc API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc
type CreateLXCResponse TaskBaseResponse
