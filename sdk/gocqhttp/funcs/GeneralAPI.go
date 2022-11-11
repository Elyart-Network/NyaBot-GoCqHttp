package funcs

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/data"
	"github.com/gorilla/websocket"
)

// SendPrivateMsg Endpoint send_private_msg
func SendPrivateMsg(c *websocket.Conn, UserId int, GroupId int, Message string, AutoEscape bool) {
	Data := data.SendPrivateMsgData{
		UserId:     UserId,
		GroupId:    GroupId,
		Message:    Message,
		AutoEscape: AutoEscape,
	}
	WsSend(c, "send_private_msg", Data)
}
