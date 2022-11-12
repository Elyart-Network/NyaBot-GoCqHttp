package EventFlow

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqdata"
	"github.com/gin-gonic/gin"
	"io"
)

func Receive(c *gin.Context) {
	context, _ := io.ReadAll(c.Request.Body)
	Process(gcqdata.Handler, context)
}
