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
			go log.Println(string(messageData))
			decode(string(messageData))
		default:
		}
	}
}
