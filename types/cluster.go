package types

// ClusterNextIdResponse maps to the GET /cluster/nextid API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/nextid
type ClusterNextIdResponse struct {
	VMID string `json:"data"`
}

// ClusterTasksResponse maps to the GET /cluster/tasks API response
// from Proxmox. See:https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/lxc/{vmid}/clone
// https://pve.proxmox.com/pve-docs/api-viewer/#/cluster/tasks
type ClusterTasksResponse struct {
	Data []ClusterTasksData `json:"data"`
}

type ClusterTasksData struct {
	UPID string `json:"upid"`
}
