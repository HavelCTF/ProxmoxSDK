package types

type VersionConsole string

const (
	Applet  VersionConsole = "applet"
	VV      VersionConsole = "vv"
	HTML5   VersionConsole = "html5"
	XTERMJS VersionConsole = "xtermjs"
)

// VersionResponse maps to the GET /version API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/version
type VersionResponse struct {
	Data VersionData `json:"data"`
}

type VersionData struct {
	Version string         `json:"version"`
	Release string         `json:"release"`
	RepoID  string         `json:"repoid"`
	Console VersionConsole `json:"console"`
}
