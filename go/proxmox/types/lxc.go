package types

type ContainerStatus string
type LXCArch string

const (
	Stopped ContainerStatus = "stopped"
	Running ContainerStatus = "running"
)

const (
	Amd64   LXCArch = "amd64"
	I386    LXCArch = "i386"
	Arm64   LXCArch = "arm64"
	Armhf   LXCArch = "armhf"
	Riscv32 LXCArch = "riscv32"
	Riscv64 LXCArch = "riscv64"
)

type GetLXCResponse struct {
	LXCs []LXCInfo `json:"data"`
}

type LXCInfo struct {
	Status ContainerStatus `json:"status"`
	VMID   int             `json:"vmid"`

	Name               string  `json:"name,omitempty"`
	Tags               string  `json:"tags,omitempty"`
	Template           bool    `json:"template,omitempty"`
	Uptime             int64   `json:"uptime,omitempty"`
	CPU                float32 `json:"cpu,omitempty"`
	CPUS               float32 `json:"cpus,omitempty"`
	Disk               uint    `json:"disk,omitempty"`
	DiskRead           int     `json:"diskread,omitempty"`
	DiskWrite          int     `json:"diskwrite,omitempty"`
	MaxDisk            int     `json:"maxdisk,omitempty"`
	Lock               string  `json:"lock,omitempty"`
	Mem                int     `json:"mem,omitempty"`
	MaxMem             int     `json:"maxmem,omitempty"`
	MaxSwap            int     `json:"maxswap,omitempty"`
	NetIn              int     `json:"netin,omitempty"`
	NetOut             int     `json:"netout,omitempty"`
	PressureCPUSome    float32 `json:"pressurecpusome,omitempty"`
	PressureIOFull     float32 `json:"pressureiofull,omitempty"`
	PressureIOSome     float32 `json:"pressureiosome,omitempty"`
	PressureMemoryFull float32 `json:"pressurememoryfull,omitempty"`
	PressureMemorySome float32 `json:"pressurememorysome,omitempty"`
}

type LXC struct {
	Node       string `json:"node"`
	OSTemplate string `json:"ostemplate"`
	VMID       int    `json:"vmid"`

	Arch LXCArch `json:"arch" default:"amd64"`
}
