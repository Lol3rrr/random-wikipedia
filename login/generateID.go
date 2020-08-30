package login

import (
	"crypto/sha256"
	"encoding/base64"
)

func generateID(email string) string {
	hasher := sha256.New()
	hasher.Write([]byte(email))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
