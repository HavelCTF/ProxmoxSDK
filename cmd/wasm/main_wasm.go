package main

import (
	"fmt"
	"syscall/js"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/cluster"
	"github.com/HavelCTF/ProxmoxSDK/internal/nodes"
	"github.com/HavelCTF/ProxmoxSDK/internal/nodes/tasks"
	"github.com/HavelCTF/ProxmoxSDK/internal/version"
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

func asyncWrapperArgs(fn func(args []js.Value) (any, error)) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		handler := js.FuncOf(func(this js.Value, promiseArgs []js.Value) any {
			resolve := promiseArgs[0]
			reject := promiseArgs[1]

			go func() {
				res, err := fn(args)
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
				return map[string]any{"VMID": res.VMID}, nil
			}),
			"GetTasks": asyncWrapper(func() (any, error) {
				res, err := clusterService.GetTasks()
				if err != nil {
					return nil, err
				}

				jsData := make([]any, len(res.Data))
				for i, task := range res.Data {
					jsData[i] = map[string]any{
						"UPID": task.UPID,
					}
				}

				return map[string]any{"Data": jsData}, nil
			}),
		},
		"version": map[string]any{
			"GetVersion": asyncWrapper(func() (any, error) {
				res, err := version.GetVersion(proxmoxClient)
				if err != nil {
					return nil, err
				}

				return map[string]any{
					"Data": map[string]any{
						"Version": res.Data.Version,
						"Release": res.Data.Release,
						"RepoID":  res.Data.RepoID,
						"Console": string(res.Data.Console),
					},
				}, nil
			}),
		},
		"nodes": map[string]any{
			"GetNodes": asyncWrapper(func() (any, error) {
				res, err := nodes.GetNodes(proxmoxClient)
				if err != nil {
					return nil, err
				}

				jsData := make([]any, len(res.Data))
				for i, n := range res.Data {
					jsData[i] = map[string]any{
						"Node":           n.Node,
						"Status":         string(n.Status),
						"Uptime":         n.Uptime,
						"SslFingerprint": n.SslFingerprint,
						"CPU":            n.CPU,
						"Level":          n.Level,
						"MaxCPU":         n.MaxCPU,
						"MaxMEM":         n.MaxMEM,
						"MEM":            n.MEM,
					}
				}

				return map[string]any{"Data": jsData}, nil
			}),
		},
		"tasks": map[string]any{
			"GetTaskStatus": asyncWrapperArgs(func(args []js.Value) (any, error) {
				node := args[0].String()
				upid := args[1].String()

				taskSvc := tasks.New(tasks.TaskContext{
					C:    proxmoxClient,
					Node: node,
					UPID: upid,
				})

				res, err := taskSvc.GetTaskStatus()
				if err != nil {
					return nil, err
				}

				return map[string]any{
					"Data": map[string]any{
						"ID":         res.Data.ID,
						"Node":       res.Data.Node,
						"PID":        res.Data.PID,
						"PStart":     res.Data.PStart,
						"StartTime":  res.Data.StartTime,
						"Type":       res.Data.Type,
						"UPID":       res.Data.UPID,
						"User":       res.Data.User,
						"Status":     res.Data.Status,
						"ExitStatus": res.Data.ExitStatus,
					},
				}, nil
			}),
			"DeleteTask": asyncWrapperArgs(func(args []js.Value) (any, error) {
				node := args[0].String()
				upid := args[1].String()

				taskSvc := tasks.New(tasks.TaskContext{
					C:    proxmoxClient,
					Node: node,
					UPID: upid,
				})

				res, err := taskSvc.DeleteTask()
				if err != nil {
					return nil, err
				}

				return map[string]any{"Data": res.Data}, nil
			}),
		},
	}))

	fmt.Println("[WASM] ProxmoxWASM registered")
	select {}
}
