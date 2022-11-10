package funcs

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/data"
	"NyaBot-GoCqHttp/sdk/gocqhttp/mods/EchoGen"
	"encoding/json"
)

func WsSend(Action string, Params struct{}) string {
	Echo := EchoGen.Generate()
	WsSendData := data.WsSendData{Action: Action, Params: Params, Echo: Echo}
	byteSlice, _ := json.MarshalIndent(WsSendData, "", "  ")
	return string(byteSlice)
}
