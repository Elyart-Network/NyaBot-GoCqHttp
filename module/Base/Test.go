package Base

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqapi"
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqdata"
	"encoding/json"
)

func test() {
	// Test
	Resp := gcqapi.GetLoginInfo()
	Struct := gcqdata.GetLoginInfoResp{}
	json.Unmarshal(Resp, &Struct)
	println(Struct.Nickname)
}
