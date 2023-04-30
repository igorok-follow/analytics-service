package helpers

import (
	"encoding/base64"
	"github.com/gorilla/securecookie"
)

func GenerateString(len int) string {
	key := securecookie.GenerateRandomKey(len)
	str := base64.StdEncoding.EncodeToString(key)

	return str
}
