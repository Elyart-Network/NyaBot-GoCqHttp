package EventFlow

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/data"
)

func Process(handler data.CqHandler, data interface{}) {
	context := string(data.([]byte))
	handler(context)
}
