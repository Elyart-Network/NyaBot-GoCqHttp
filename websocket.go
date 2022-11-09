package main

import (
	"github.com/gorilla/websocket" //这里使用的是 gorilla 的 websocket 库
	"log"
)

func main() {
	dialer := websocket.DefaultDialer
	connect, _, err := dialer.Dial("ws://127.0.0.1:4000/", nil)
	if nil != err {
		log.Println(err)
		return
	}
	defer connect.Close()
	worker(connect)
}

func tickWriter(connect *websocket.Conn) {
	err := connect.WriteMessage(websocket.TextMessage, []byte("from client to server"))
	if nil != err {
		log.Println(err)
		return
	}
}
