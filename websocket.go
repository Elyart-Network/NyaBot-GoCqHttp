package main

import (
	"fmt"
	"github.com/gorilla/websocket" //这里使用的是 gorilla 的 websocket 库
	"log"
	"time"
)

func loader(connect *websocket.Conn) {

}

func main() {
	dialer := websocket.DefaultDialer
	connect, _, err := dialer.Dial("ws://127.0.0.1:4000/", nil)
	if nil != err {
		log.Println(err)
		return
	}
	defer connect.Close()
	go loader(connect)

	for {
		messageType, messageData, err := connect.ReadMessage()
		if nil != err {
			log.Println(err)
			break
		}
		switch messageType {
		case websocket.TextMessage:
			fmt.Println(string(messageData))
		//case websocket.CloseMessage:
		//case websocket.PingMessage:
		//case websocket.PongMessage:
		default:
		}
	}
}

func tickWriter(connect *websocket.Conn) {
	for {
		//向客户端发送类型为文本的数据
		err := connect.WriteMessage(websocket.TextMessage, []byte("from client to server"))
		if nil != err {
			log.Println(err)
			break
		}
		//休息一秒
		time.Sleep(time.Second)
	}
}
