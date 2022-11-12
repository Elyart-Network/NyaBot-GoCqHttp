package GcqServer

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqdata"
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqmods/EventFlow"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func Init(handler gcqdata.CqHandler) {
	gcqdata.Handler = handler
}

func SetCqHost(cqHttpHost string) {
	gcqdata.CqHttpHost = cqHttpHost
}

func SetRunPort(serverPort string) {
	gcqdata.ServerPort = serverPort
}

func SetDebug(debugMode bool) {
	gcqdata.DebugMode = debugMode
}

func SetFileLog(fileLogger bool) {
	gcqdata.FileLogger = fileLogger
}

func SetEchoHash(echoHash string) {
	gcqdata.EchoHash = echoHash
}

func RunServe() {
	if gcqdata.FileLogger == true {
		gin.DisableConsoleColor()
		file, _ := os.Create("app.log")
		gin.DefaultWriter = io.MultiWriter(file)
	}
	if gcqdata.DebugMode == false {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	server.GET("/health", Health)
	server.POST("/data", EventFlow.Receive)
	err := http.ListenAndServe(":"+gcqdata.ServerPort, server)
	if err != nil {
		log.Panicln(err)
	}
}
