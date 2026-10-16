// SPDX-License-Identifier: MIT
// Copyright (c) 2026 HavelCTF

package nodes

import "github.com/HavelCTF/ProxmoxSDK/types"

type CreateLXCFinalData struct {
	Node string `url:"node"`
	types.CreateLXCData
}
