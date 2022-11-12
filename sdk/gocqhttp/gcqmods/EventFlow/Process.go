package EventFlow

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqdata"
)

func Process(handler gcqdata.CqHandler, context interface{}) {
	result := string(context.([]byte))
	handler(result)
}
