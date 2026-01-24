package types

type VersionConsole string

const (
	Applet  VersionConsole = "applet"
	VV      VersionConsole = "vv"
	HTML5   VersionConsole = "html5"
	XTERMJS VersionConsole = "xtermjs"
)

// VersionResponse represents the response from the Proxmox /version endpoint.
type VersionResponse struct {
	Data VersionData `json:"data"`
}

// VersionData contains Proxmox version information.
type VersionData struct {
	Version string         `json:"version"`
	Release string         `json:"release"`
	RepoID  string         `json:"repoid"`
	Console VersionConsole `json:"console"`
}
