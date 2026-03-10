package main

import (
	"fmt"
	"syscall/js"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/cluster"
)

type NextIdResponse struct {
	ID string `js:"id"`
}

var proxmoxClient *client.Client
var clusterService *cluster.ClusterService

// Wrapper pour transformer les appels Go synchrones en Promises JavaScript
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

// Initialise le client Go avec host, token et uuid
func initClient(this js.Value, args []js.Value) any {
	host := args[0].String()
	token := args[1].String()
	uuid := args[2].String()

	fmt.Println("initClient called with host:", host)

	proxmoxClient = client.NewClient(host, token, uuid)
	clusterService = cluster.New(proxmoxClient)

	return true
}

// Enveloppe la fonction GetNextId
func GetNextId(this js.Value, args []js.Value) any {
	res, err := clusterService.GetNextId()
	if err != nil {
		panic(err)
	}

	fmt.Println("VMID:", res.VMID)

	return js.ValueOf(NextIdResponse{ID: res.VMID})
}

func main() {
	fmt.Println("🚀 [GO] Démarrage du binaire WASM...")
	js.Global().Set("ProxmoxWASM", js.ValueOf(map[string]any{
		"initClient": js.FuncOf(initClient),
		"cluster": map[string]any{
			"GetNextId": js.FuncOf(GetNextId),
		},
	}))

	fmt.Println("✅ [GO] Objet 'ProxmoxWASM' injecté avec succès dans globalThis !")
	select {}
}
