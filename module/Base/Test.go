package Base

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqapi"
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqdata"
)

func test() {
	// Test
	Struct := gcqdata.GetLoginInfoResp{}
	gcqapi.GetLoginInfo(Struct)
	println(Struct.Nickname)
}
