package GcqServer

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/data"
	"NyaBot-GoCqHttp/sdk/gocqhttp/mods/EventFlow"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func Init(handler data.CqHandler) {
	data.Handler = handler
}

func SetCqHost(cqHttpHost string) {
	data.CqHttpHost = cqHttpHost
}

func SetRunPort(serverPort string) {
	data.ServerPort = serverPort
}

func SetDebug(debugMode bool) {
	data.DebugMode = debugMode
}

func SetFileLog(fileLogger bool) {
	data.FileLogger = fileLogger
}

func RunServe() {
	if data.FileLogger == true {
		gin.DisableConsoleColor()
		file, _ := os.Create("app.log")
		gin.DefaultWriter = io.MultiWriter(file)
	}
	if data.DebugMode == false {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	server.GET("/health", Health)
	server.POST("/data", EventFlow.Receive)
	err := http.ListenAndServe(":"+data.ServerPort, server)
	if err != nil {
		log.Panicln(err)
	}
}
