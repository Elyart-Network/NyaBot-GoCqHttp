package EchoGen

import (
	"NyaBot-GoCqHttp/sdk/gocqhttp/gcqdata"
	"golang.org/x/crypto/bcrypt"
)

func Generate() string {
	// Generate a hash from password string
	encodedEcho, _ := bcrypt.GenerateFromPassword([]byte(gcqdata.EchoHash), 3)
	return string(encodedEcho)
}
