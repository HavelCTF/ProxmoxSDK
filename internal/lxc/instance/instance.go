package instance

import (
	"context"
	"fmt"
	"time"

	lxccontext "github.com/HavelCTF/ProxmoxSDK/internal/context/lxc"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

type LXCInstance struct {
	ctx lxccontext.LXCContext
}

func New(ctx lxccontext.LXCContext, vmid int) *LXCInstance {
	instanceCtx := ctx
	instanceCtx.VMID = &vmid
	return &LXCInstance{
		ctx: instanceCtx,
	}
}

func (i *LXCInstance) Delete() (*types.LXCResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	return http.DoRequest[types.LXCResponse](ctx, i.ctx.Client,
		http.RequestContent{
			Method:   "DELETE",
			Endpoint: fmt.Sprintf("/nodes/%s/lxc/%d", i.ctx.Node, *i.ctx.VMID),
		},
	)
}
