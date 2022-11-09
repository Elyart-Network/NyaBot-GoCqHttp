package main

import (
	"NyaBot-GoCqHttp/actions/Command"
	"github.com/gorilla/websocket"
	"log"
)

const wsHost = "ws://127.0.0.1:4000/"

func main() {
	dialer := websocket.DefaultDialer
	connect, _, err := dialer.Dial(wsHost, nil)
	if nil != err {
		log.Println(err)
		return
	}
	defer connect.Close()
	for {
		messageType, messageData, err := connect.ReadMessage()
		if nil != err {
			log.Println(err)
			break
		}
		switch messageType {
		case websocket.TextMessage:
			log.Println(string(messageData))
			go Command.Loader()
		default:
		}
	}
}

//Example
//func tickWriter(connect *websocket.Conn) {
//	err := connect.WriteMessage(websocket.TextMessage, []byte("from client to server"))
//	if nil != err {
//		log.Println(err)
//		return
//	}
//}
