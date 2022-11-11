package EventFlow

import (
	"github.com/gin-gonic/gin"
	"io"
)

func Receive(c *gin.Context) {
	data, _ := io.ReadAll(c.Request.Body)
	Process(data)
}
