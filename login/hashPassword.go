package login

import (
	"crypto/sha512"
	"encoding/base64"
)

func hashPassword(password string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
