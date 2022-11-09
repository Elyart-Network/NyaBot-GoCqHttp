package EchoGen

import "golang.org/x/crypto/bcrypt"

const Echo = "DtLoKo2fDNsqRVmsibBoMjwvRwMTWN"

func Generate() string {
	encodedEcho, _ := bcrypt.GenerateFromPassword([]byte(Echo), 3)
	return string(encodedEcho)
}
