package types

// StorageContentResponse maps to the GET /nodes/{node}/storage/{storage}/content API response
// from Proxmox. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/content
type StorageContentResponse struct {
	Data []StorageContentData `json:"data"`
}

type StorageContentData struct {
	VolID   string `json:"volid"`
	Content string `json:"content"`
	Format  string `json:"format,omitempty"`
	Size    int    `json:"size,omitempty"`
	Used    int    `json:"used,omitempty"`
	CTime   int    `json:"ctime,omitempty"`
	VMID    int    `json:"vmid,omitempty"`
	Notes   string `json:"notes,omitempty"`
	Parent  string `json:"parent,omitempty"`
	Verify  any    `json:"verification,omitempty"`
}

// UploadResponse maps to the POST /nodes/{node}/storage/{storage}/upload API response
// from Proxmox. The endpoint is asynchronous and returns a UPID. See:
// https://pve.proxmox.com/pve-docs/api-viewer/#/nodes/{node}/storage/{storage}/upload
type UploadResponse TaskBaseResponse
