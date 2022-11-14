package Module

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqapi"
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqdata"
	"encoding/json"
)

func Handler(context interface{}) {
	Msg := gcqapi.GetLoginInfo()
	Struct := gcqdata.GetLoginInfoResp{}
	json.Unmarshal(Msg, &Struct)
	println(Struct.Nickname)
}
