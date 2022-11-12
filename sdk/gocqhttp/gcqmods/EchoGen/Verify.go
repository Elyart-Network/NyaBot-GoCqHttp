package EchoGen

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqdata"
	"golang.org/x/crypto/bcrypt"
)

func Verify(echo string) bool {
	echoOri := gcqdata.EchoHash
	err := bcrypt.CompareHashAndPassword([]byte(echo), []byte(echoOri))
	if err == nil {
		return true
	}
	return false
}
