package GcqServer

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/mods/EventFlow"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

var CqHttpHost = "http://127.0.0.1:4000/"
var serverPort = "3000"
var debugMode = false
var fileLogger = false

func SetCqHost(cqHttpHost string) {
	CqHttpHost = cqHttpHost
}

func SetRunPort(ServerPort string) {
	serverPort = ServerPort
}

func SetDebug(DebugMode bool) {
	debugMode = DebugMode
}

func SetFileLog(FileLogger bool) {
	fileLogger = FileLogger
}

func RunServe() {
	if fileLogger == true {
		gin.DisableConsoleColor()
		file, _ := os.Create("app.log")
		gin.DefaultWriter = io.MultiWriter(file)
	}
	if debugMode == false {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	server.GET("/health", Health)
	server.POST("/data", EventFlow.Receive)
	err := http.ListenAndServe(":"+serverPort, server)
	if err != nil {
		log.Panicln(err)
	}
}
