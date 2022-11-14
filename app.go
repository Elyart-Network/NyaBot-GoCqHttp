package main

import (
	"NyaBot-GoCqHttp/module"
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqmods/GcqServer"
)

func init() {
	// Initialize GcqServer
	GcqServer.SetCqHost("http://127.0.0.1:4000/")
	GcqServer.SetRunPort("3000")
	GcqServer.SetDebug(true)
	GcqServer.SetFileLog(false)
}

func main() {
	// Start GcqServer
	GcqServer.Init(module.Handler)
	GcqServer.RunServe()
}
