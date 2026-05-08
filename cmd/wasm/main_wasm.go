package main

import (
	"bytes"
	"context"
	"fmt"
	"syscall/js"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/cluster"
	"github.com/HavelCTF/ProxmoxSDK/internal/nodes"
	"github.com/HavelCTF/ProxmoxSDK/internal/nodes/lxc"
	"github.com/HavelCTF/ProxmoxSDK/internal/nodes/storage"
	"github.com/HavelCTF/ProxmoxSDK/internal/nodes/tasks"
	"github.com/HavelCTF/ProxmoxSDK/internal/version"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

var proxmoxClient *client.Client

func asyncWrapper(fn func() (any, error)) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		handler := js.FuncOf(func(this js.Value, promiseArgs []js.Value) any {
			resolve := promiseArgs[0]
			reject := promiseArgs[1]

			go func() {
				res, err := fn()
				if err != nil {
					errObj := js.Global().Get("Error").New(err.Error())
					reject.Invoke(errObj)
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
					errObj := js.Global().Get("Error").New(err.Error())
					reject.Invoke(errObj)
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

	fmt.Println("[WASM] client initialized for", host)
	return true
}

func main() {
	js.Global().Set("ProxmoxWASM", js.ValueOf(map[string]any{
		"initClient": js.FuncOf(initClient),
		"cluster": map[string]any{
			"GetNextId": asyncWrapper(func() (any, error) {

				clusterSvc := cluster.New(proxmoxClient)

				res, err := clusterSvc.GetNextId()
				if err != nil {
					return nil, err
				}
				return map[string]any{"VMID": res.VMID}, nil
			}),
			"GetTasks": asyncWrapper(func() (any, error) {

				clusterSvc := cluster.New(proxmoxClient)

				res, err := clusterSvc.GetTasks()
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
			"GetTasks": asyncWrapperArgs(func(args []js.Value) (any, error) {
				nodeName := args[0].String()
				nodeSvc := nodes.New(proxmoxClient, nodeName)

				res, err := nodeSvc.GetTasks()
				if err != nil {
					return nil, err
				}

				jsData := make([]any, len(res.Data))
				for i, task := range res.Data {
					jsData[i] = map[string]any{
						"UPID":      task.UPID,
						"Node":      task.Node,
						"PID":       task.PID,
						"PStart":    task.PStart,
						"StartTime": task.StartTime,
						"Type":      task.Type,
						"User":      task.User,
						"Status":    task.Status,
					}
				}

				return map[string]any{"Data": jsData}, nil
			}),
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
			"Wait": asyncWrapperArgs(func(args []js.Value) (any, error) {
				node := args[0].String()
				upid := args[1].String()

				pollIntervalMs := 1000
				timeoutMs := 0
				if len(args) > 2 && !args[2].IsUndefined() && !args[2].IsNull() {
					if v := args[2].Get("PollIntervalMs"); !v.IsUndefined() && !v.IsNull() {
						pollIntervalMs = v.Int()
					}
					if v := args[2].Get("TimeoutMs"); !v.IsUndefined() && !v.IsNull() {
						timeoutMs = v.Int()
					}
				}

				ctx := context.Background()
				if timeoutMs > 0 {
					var cancel context.CancelFunc
					ctx, cancel = context.WithTimeout(ctx, time.Duration(timeoutMs)*time.Millisecond)
					defer cancel()
				}

				taskSvc := tasks.New(tasks.TaskContext{
					C:    proxmoxClient,
					Node: node,
					UPID: upid,
				})

				res, err := taskSvc.Wait(ctx, time.Duration(pollIntervalMs)*time.Millisecond)
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
		},
		"lxc": map[string]any{
			"StartLXC": asyncWrapperArgs(func(args []js.Value) (any, error) {
				node := args[0].String()
				vmid := args[1].Int()

				svc := lxc.New(lxc.LXCContext{
					Client: proxmoxClient,
					Node:   node,
					VMID:   vmid,
				})

				res, err := svc.Status().StartLXC()
				if err != nil {
					return nil, err
				}

				return map[string]any{"Data": res.UPID}, nil
			}),
			"StopLXC": asyncWrapperArgs(func(args []js.Value) (any, error) {
				node := args[0].String()
				vmid := args[1].Int()

				svc := lxc.New(lxc.LXCContext{
					Client: proxmoxClient,
					Node:   node,
					VMID:   vmid,
				})

				res, err := svc.Status().StopLXC()
				if err != nil {
					return nil, err
				}

				return map[string]any{"Data": res.UPID}, nil
			}),
			"DeleteLXC": asyncWrapperArgs(func(args []js.Value) (any, error) {
				node := args[0].String()
				vmid := args[1].Int()

				svc := lxc.New(lxc.LXCContext{
					Client: proxmoxClient,
					Node:   node,
					VMID:   vmid,
				})

				res, err := svc.DeleteLXC()
				if err != nil {
					return nil, err
				}

				return map[string]any{"Data": res.UPID}, nil
			}),
			"CloneLXC": asyncWrapperArgs(func(args []js.Value) (any, error) {
				node := args[0].String()
				vmid := args[1].Int()
				dataObj := args[2]

				newId := dataObj.Get("NewId").Int()
				var target *string
				if t := dataObj.Get("Target"); !t.IsUndefined() && !t.IsNull() {
					targetStr := t.String()
					target = &targetStr
				}

				svc := lxc.New(lxc.LXCContext{
					Client: proxmoxClient,
					Node:   node,
					VMID:   vmid,
				})

				res, err := svc.CloneLXC(types.CloneLXCData{
					NewId:  newId,
					Target: target,
				})
				if err != nil {
					return nil, err
				}

				return map[string]any{"Data": res.UPID}, nil
			}),
			"GetLXCs": asyncWrapperArgs(func(args []js.Value) (any, error) {
				nodeName := args[0].String()
				nodeSvc := nodes.New(proxmoxClient, nodeName)

				res, err := nodeSvc.GetLXCs()
				if err != nil {
					return nil, err
				}

				jsData := make([]any, len(res.LXCs))
				for i, lxcData := range res.LXCs {
					jsData[i] = map[string]any{
						"Status":             string(lxcData.Status),
						"VMID":               lxcData.VMID,
						"CPU":                lxcData.CPU,
						"CPUS":               lxcData.CPUS,
						"Disk":               lxcData.Disk,
						"DiskRead":           lxcData.DiskRead,
						"DiskWrite":          lxcData.DiskWrite,
						"Lock":               lxcData.Lock,
						"MaxDisk":            lxcData.MaxDisk,
						"MaxMem":             lxcData.MaxMem,
						"MaxSwap":            lxcData.MaxSwap,
						"Mem":                lxcData.Mem,
						"Name":               lxcData.Name,
						"NetIn":              lxcData.NetIn,
						"NetOut":             lxcData.NetOut,
						"PressureCPUSome":    lxcData.PressureCPUSome,
						"PressureIOFull":     lxcData.PressureIOFull,
						"PressureIOSome":     lxcData.PressureIOSome,
						"PressureMemoryFull": lxcData.PressureMemoryFull,
						"PressureMemorySome": lxcData.PressureMemorySome,
						"Tags":               lxcData.Tags,
						"Template":           lxcData.Template,
						"Uptime":             lxcData.Uptime,
					}
				}

				return map[string]any{"LXCs": jsData}, nil
			}),
			"Get": asyncWrapperArgs(func(args []js.Value) (any, error) {
				node := args[0].String()
				vmid := args[1].Int()

				svc := lxc.New(lxc.LXCContext{
					Client: proxmoxClient,
					Node:   node,
					VMID:   vmid,
				})

				res, err := svc.Get()
				if err != nil {
					return nil, err
				}

				return map[string]any{
					"Data": map[string]any{
						"Name":      res.Data.Name,
						"Status":    string(res.Data.Status),
						"VMID":      res.Data.VMID,
						"Uptime":    res.Data.Uptime,
						"CPUs":      res.Data.CPUs,
						"CPU":       res.Data.CPU,
						"Mem":       res.Data.Mem,
						"MaxMem":    res.Data.MaxMem,
						"Disk":      res.Data.Disk,
						"MaxDisk":   res.Data.MaxDisk,
						"Swap":      res.Data.Swap,
						"MaxSwap":   res.Data.MaxSwap,
						"NetIn":     res.Data.NetIn,
						"NetOut":    res.Data.NetOut,
						"DiskRead":  res.Data.DiskRead,
						"DiskWrite": res.Data.DiskWrite,
						"Lock":      res.Data.Lock,
						"Tags":      res.Data.Tags,
						"Type":      res.Data.Type,
					},
				}, nil
			}),
			"PostLXC": asyncWrapperArgs(func(args []js.Value) (any, error) {
				nodeName := args[0].String()
				dataObj := args[1]

				nodeSvc := nodes.New(proxmoxClient, nodeName)

				osTemplate := dataObj.Get("OSTemplate").String()
				vmid := dataObj.Get("VMID").Int()

				var features *types.LXCFeatures
				if f := dataObj.Get("Features"); !f.IsUndefined() && !f.IsNull() {
					features = &types.LXCFeatures{
						Nesting: f.Get("Nesting").Bool(),
					}
				}

				res, err := nodeSvc.PostLXC(types.CreateLXCData{
					OSTemplate: osTemplate,
					VMID:       vmid,
					Features:   features,
				})
				if err != nil {
					return nil, err
				}

				return map[string]any{"Data": res.UPID}, nil
			}),
		},
		"storage": map[string]any{
			"GetContent": asyncWrapperArgs(func(args []js.Value) (any, error) {
				node := args[0].String()
				storageName := args[1].String()
				contentType := ""
				if len(args) > 2 && !args[2].IsUndefined() && !args[2].IsNull() {
					contentType = args[2].String()
				}

				svc := storage.New(storage.StorageContext{
					Client:  proxmoxClient,
					Node:    node,
					Storage: storageName,
				})

				res, err := svc.GetContent(contentType)
				if err != nil {
					return nil, err
				}

				jsData := make([]any, len(res.Data))
				for i, item := range res.Data {
					jsData[i] = map[string]any{
						"VolID":   item.VolID,
						"Content": item.Content,
						"Format":  item.Format,
						"Size":    item.Size,
						"Used":    item.Used,
						"CTime":   item.CTime,
						"VMID":    item.VMID,
						"Notes":   item.Notes,
						"Parent":  item.Parent,
					}
				}

				return map[string]any{"Data": jsData}, nil
			}),
			"UploadTemplate": asyncWrapperArgs(func(args []js.Value) (any, error) {
				node := args[0].String()
				storageName := args[1].String()
				filename := args[2].String()
				bodyJS := args[3]

				// JS passes the payload as a Uint8Array; copy it into a Go []byte
				// and wrap it in bytes.NewReader so the multipart pipe can consume it.
				length := bodyJS.Get("byteLength").Int()
				buf := make([]byte, length)
				js.CopyBytesToGo(buf, bodyJS)

				svc := storage.New(storage.StorageContext{
					Client:  proxmoxClient,
					Node:    node,
					Storage: storageName,
				})

				res, err := svc.UploadTemplate(filename, bytes.NewReader(buf), int64(length))
				if err != nil {
					return nil, err
				}

				return map[string]any{"Data": res.UPID}, nil
			}),
		},
	}))

	fmt.Println("[WASM] ProxmoxWASM registered")
	select {}
}
