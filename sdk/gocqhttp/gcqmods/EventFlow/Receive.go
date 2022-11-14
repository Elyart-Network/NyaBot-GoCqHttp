package EventFlow

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqdata"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func Receive(c *gin.Context) {
	context, _ := io.ReadAll(c.Request.Body)
	if gcqdata.DebugMode == true {
		log.Println("Receive: " + string(context))
	}
	Process(gcqdata.Handler, context)
}
