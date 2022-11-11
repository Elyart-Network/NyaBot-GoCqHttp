package Module

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/data"
)

func Handler(data data.GeneralEvent) {
	println(data.PostType)
}
