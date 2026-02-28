package types

type ClusterNextIdResponse struct {
	VMID int `json:"data"`
}

type ClusterTasksResponse struct {
	Data []ClusterTasksData `json:"data"`
}

type ClusterTasksData struct {
	UPID string `json:"upid"`
}
