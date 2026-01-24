package types

// VersionResponse represents the response from the Proxmox /version endpoint.
type VersionResponse struct {
	Data VersionData `json:"data"`
}

// VersionData contains Proxmox version information.
type VersionData struct {
	Version string `json:"version"`
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
}
