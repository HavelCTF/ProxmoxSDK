package main

import (
	"fmt"
	"syscall/js"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/cluster"
)

var proxmoxClient *client.Client
var clusterService *cluster.ClusterService

func asyncWrapper(fn func() (any, error)) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		handler := js.FuncOf(func(this js.Value, promiseArgs []js.Value) any {
			resolve := promiseArgs[0]
			reject := promiseArgs[1]

			go func() {
				res, err := fn()
				if err != nil {
					reject.Invoke(err.Error())
					return
				}
				resolve.Invoke(res)
			}()
			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}

func initClient(this js.Value, args []js.Value) any {
	host := args[0].String()
	token := args[1].String()
	uuid := args[2].String()

	proxmoxClient = client.NewClient(host, token, uuid)
	clusterService = cluster.New(proxmoxClient)

	fmt.Println("[WASM] client initialized for", host)
	return true
}

func main() {
	js.Global().Set("ProxmoxWASM", js.ValueOf(map[string]any{
		"initClient": js.FuncOf(initClient),
		"cluster": map[string]any{
			"GetNextId": asyncWrapper(func() (any, error) {
				res, err := clusterService.GetNextId()
				if err != nil {
					return nil, err
				}
				return map[string]any{"id": res.VMID}, nil
			}),
		},
	}))

	fmt.Println("[WASM] ProxmoxWASM registered")
	select {}
}
