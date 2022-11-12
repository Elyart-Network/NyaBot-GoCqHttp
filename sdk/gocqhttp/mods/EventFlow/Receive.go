package EventFlow

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/data"
	"github.com/gin-gonic/gin"
	"io"
)

func Receive(c *gin.Context) {
	context, _ := io.ReadAll(c.Request.Body)
	Process(data.Handler, context)
}
