package main

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/mods/GcqServer"
)

func init() {
	GcqServer.SetCqHost("http://127.0.0.1:4000/")
	GcqServer.SetRunPort("3000")
	GcqServer.SetDebug(true)
	GcqServer.SetFileLog(false)
}

func main() {
	GcqServer.RunServe()
}
