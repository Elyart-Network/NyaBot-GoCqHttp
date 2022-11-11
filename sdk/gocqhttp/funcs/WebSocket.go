package funcs

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/data"
	"NyaBot-GoCqHttp/sdk/gocqhttp/mods/EchoGen"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

func WsSend(c *websocket.Conn, Action string, Params interface{}) {
	Echo := EchoGen.Generate()
	Data := data.WsSendData{
		Action: Action,
		Params: Params,
		Echo:   Echo,
	}
	byteSlice, _ := json.MarshalIndent(Data, "", "")
	println(string(byteSlice))
	err := c.WriteMessage(websocket.TextMessage, byteSlice)
	if nil != err {
		log.Println(err)
		return
	}
}
