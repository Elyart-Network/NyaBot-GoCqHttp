package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

func decode(message string) {
	reading := make(map[string]interface{})
	err := json.Unmarshal([]byte(message), &reading)
	if err != nil {
		log.Println(err)
	}
}

func worker(connect *websocket.Conn) {
	for {
		messageType, messageData, err := connect.ReadMessage()
		if nil != err {
			log.Println(err)
			break
		}
		switch messageType {
		case websocket.TextMessage:
			decode(string(messageData))
			go log.Println(string(messageData))
		default:
		}
	}
}
