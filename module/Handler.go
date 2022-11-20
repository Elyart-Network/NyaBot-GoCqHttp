package module

import (
	"NyaBot-GoCqHttp/module/Base"
	"NyaBot-GoCqHttp/module/Common"
)

func Handler(c interface{}) {
	// Initialize all modules
	go Base.Loader(c)
	go Common.Loader(c)
}
