package types

type VersionResponse struct {
	Data VersionData `json:"data"`
}

type VersionData struct {
	Version string `json:"version"`
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
}
