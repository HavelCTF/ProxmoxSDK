package nodes

import "github.com/HavelCTF/ProxmoxSDK/types"

type CreateLXCFinalData struct {
	Node string `url:"node"`
	types.CreateLXCData
}
