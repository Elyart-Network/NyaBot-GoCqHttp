package EventFlow

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/data"
)

func Process(handler data.CqHandler, context interface{}) {
	result := string(context.([]byte))
	handler(result)
}
