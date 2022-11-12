package main

import (
	"NyaBot-GoCqHttp/actions/Module"
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqmods/GcqServer"
)

func init() {
	// Initialize GcqServer
	GcqServer.SetCqHost("http://127.0.0.1:4000/")
	GcqServer.SetRunPort("3000")
	GcqServer.SetDebug(true)
	GcqServer.SetFileLog(false)
	GcqServer.SetEchoHash("123456")
}

func handler(context interface{}) {
	// Module Register
	Module.Handler(context)
}

func main() {
	// Start GcqServer
	GcqServer.Init(handler)
	GcqServer.RunServe()
}
