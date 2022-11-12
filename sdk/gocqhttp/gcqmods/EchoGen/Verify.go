package EchoGen

import "golang.org/x/crypto/bcrypt"

func Verify(echo string) bool {
	echoOri := Echo
	err := bcrypt.CompareHashAndPassword([]byte(echo), []byte(echoOri))
	if err == nil {
		return true
	}
	return false
}
