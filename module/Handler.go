package module

import (
	"NyaBot-GoCqHttp/module/Base"
	"NyaBot-GoCqHttp/module/Common"
)

func Handler(c interface{}) {
	// Initialize all modules
	Base.Loader(c)
	Common.Loader(c)
}
