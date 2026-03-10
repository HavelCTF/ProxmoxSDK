package lxc

import "github.com/HavelCTF/ProxmoxSDK/types"

type CloneLXCFinalData struct {
	types.CloneLXCData
	Node string `url:"node"`
	VMID int    `url:"vmid"`
}
