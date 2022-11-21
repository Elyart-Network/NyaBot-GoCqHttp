package Base

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqapi"
	"log"
)

func test() {
	// Test
	Struct := gcqapi.GetLoginInfo()
	log.Println(Struct.Nickname)
}
