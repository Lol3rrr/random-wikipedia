package login

import (
	"math/rand"
	"strings"
)

var options string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$/()=?"
var optionLength int = 70

func generatePassword(length int) string {
	var result strings.Builder
	for i := 0; i < length; i++ {
		result.WriteByte(options[rand.Intn(70)])
	}
	return result.String()
}
