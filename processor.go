package main

import (
	"NyaBot-GoCqHttp/actions/Command"
	"github.com/gorilla/websocket"
	"log"
)

func worker(connect *websocket.Conn) {
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
