package types

type LXCStatus string

// type LXCArch string
// type CMode string
// type Lock string
// type OSType string
// type LXCDevMode uint16
// type LXCMount string
// type LXCMountOption string
// type DiskUnit string
// type LXCIP string

const (
	Stopped LXCStatus = "stopped"
	Running LXCStatus = "running"
)

// const (
// 	Amd64   LXCArch = "amd64"
// 	I386    LXCArch = "i386"
// 	Arm64   LXCArch = "arm64"
// 	Armhf   LXCArch = "armhf"
// 	Riscv32 LXCArch = "riscv32"
// 	Riscv64 LXCArch = "riscv64"
// )

// const (
// 	Shell   CMode = "shell"
// 	Console CMode = "console"
// 	TTY     CMode = "tty"
// )

// const (
// 	Backup         Lock = "backup"
// 	Create         Lock = "create"
// 	Destroyed      Lock = "destroyed"
// 	Disk           Lock = "disk"
// 	Fstrim         Lock = "fstrim"
// 	Migrate        Lock = "migrate"
// 	Mounted        Lock = "mounted"
// 	Rollback       Lock = "rollback"
// 	Snapshot       Lock = "snapshot"
// 	SnapshotDelete Lock = "snapshot-delete"
// )

// const (
// 	Debian    OSType = "debian"
// 	Devuan    OSType = "devuan"
// 	Ubuntu    OSType = "ubuntu"
// 	Centos    OSType = "centos"
// 	Fedora    OSType = "fedora"
// 	Opensuse  OSType = "opensuse"
// 	ArchLinux OSType = "archlinux"
// 	Alpine    OSType = "alpine"
// 	Gentoo    OSType = "gentoo"
// 	NixOS     OSType = "nixos"
// 	Unmanaged OSType = "unmanaged"
// )

// const (
// 	Mode600 LXCDevMode = 0o600
// 	Mode644 LXCDevMode = 0o644
// 	Mode660 LXCDevMode = 0o660
// 	Mode664 LXCDevMode = 0o664
// 	Mode700 LXCDevMode = 0o700
// 	Mode755 LXCDevMode = 0o755
// 	Mode770 LXCDevMode = 0o770
// 	Mode775 LXCDevMode = 0o775
// 	Mode777 LXCDevMode = 0o777
// )
// const (
// 	MountFUSE    LXCMount = "fuse"
// 	MountNFS     LXCMount = "nfs"
// 	MountCIFS    LXCMount = "cifs"
// 	MountPROC    LXCMount = "proc"
// 	MountSYSFS   LXCMount = "sysfs"
// 	MountOVERLAY LXCMount = "overlay"
// 	MountTMPFS   LXCMount = "tmpfs"
// )

// const (
// 	// RW Only
// 	LazyTime     LXCMountOption = "lazytime"
// 	MountDiscard LXCMountOption = "discard"
// 	MountNoAtime LXCMountOption = "noatime"

// 	// RW + RO
// 	MountNoSuid LXCMountOption = "nosuid"
// 	MountNoDev  LXCMountOption = "nodev"
// 	MountNoExec LXCMountOption = "noexec"
// )

// const (
// 	KB DiskUnit = "K"
// 	MB DiskUnit = "M"
// 	GB DiskUnit = "G"
// 	TB DiskUnit = "T"
// )

// const (
// 	//IPv4
// 	IPv4 LXCIP = "IPv4"

// 	//IPv6
// 	IPv6 LXCIP = "IPv6"
// 	Auto LXCIP = "auto"

// 	//IPv4 + IPv6
// 	CIDR   LXCIP = "CIDR"
// 	DHCP   LXCIP = "dhcp"
// 	Manual LXCIP = "manual"
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
	PressureCPUSome    float32   `json:"pressurecpusome,omitempty"`
	PressureIOFull     float32   `json:"pressureiofull,omitempty"`
	PressureIOSome     float32   `json:"pressureiosome,omitempty"`
	PressureMemoryFull float32   `json:"pressurememoryfull,omitempty"`
	PressureMemorySome float32   `json:"pressurememorysome,omitempty"`
	Tags               string    `json:"tags,omitempty"`
	Template           bool      `json:"template,omitempty"`
	Uptime             int       `json:"uptime,omitempty"`
}

// type LXC struct {
// 	Node       string
// 	OSTemplate string
// 	VMID       int

// 	Arch               LXCArch
// 	BwLimit            float32
// 	CMode              CMode
// 	Console            bool
// 	Cores              int
// 	CPULimit           float32
// 	CPUUnits           int
// 	Debug              bool
// 	Description        string
// 	EntryPoint         string
// 	Env                string
// 	Force              bool
// 	HAManaged          bool
// 	HookScript         string
// 	Hostname           string
// 	IgnoreUnpackErrors bool
// 	Lock               Lock
// 	Memory             int
// 	Nameserver         string
// 	OnBoot             bool
// 	OSType             OSType
// 	Password           string
// 	Pool               string
// 	Protection         bool
// 	Restore            bool
// 	SearchDomain       string
// 	SSHPublicKeys      string
// 	Start              bool
// 	Startup            string
// 	Storage            string
// 	Swap               int
// 	Tags               string
// 	Template           bool
// 	TimeZone           string
// 	TTY                int
// 	Unique             bool
// 	Unprivileged       bool

// 	Features *LXCFeatures
// 	RootFS   *LXCVolume
// 	Dev      *[]LXCDev
// 	Mp       *[]LXCMp
// 	//Net      *[]LXCNet TODO: Network imp when proxmox ready
// 	Unused *[]LXCUnused
// }

// type LXCFeatures struct {
// 	ForceRwSys bool
// 	Fuse       *bool
// 	KeyCTL     *bool
// 	MkNod      *bool
// 	Mount      *[]LXCMount
// 	Nesting    *bool
// }

// type LXCDev struct {
// 	Path      string
// 	DenyWrite *bool
// 	GID       *int
// 	Mode      *LXCDevMode
// 	UID       *int
// }

// type LXCMp struct {
// 	Volume LXCVolume
// 	Mp     string
// 	Backup *bool
// }

// type LXCVolume struct {
// 	Volume       string
// 	ACL          *bool
// 	MountOptions *[]LXCMountOption
// 	Quota        *bool
// 	Replicate    *bool
// 	Ro           *bool
// 	Shared       *bool
// 	Size         *DiskSize
// }

// type DiskSize struct {
// 	Value uint64
// 	Unit  DiskUnit
// }

/* TODO: Net COnfiguration For LXC
type LXCNet struct {
	Name string
	//TODO [,bridge=<bridge>]
	Firewall    bool
	GatewayIPv4 string
	GatewayIPv6 string
	HostManaged bool
	HWAddr      string
	IP          LXCIP
	IPv6        LXCIP
	LinkDown    bool
	MTU         int
	Rate        int
	Tag         int
	//TODO [,trunks=<vlanid[;vlanid...]>]
	//TODO [,type=<veth>]
}
*/

// type LXCUnused struct {
// 	Volume string
// }
