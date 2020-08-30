package general

import (
	"os"
	"strconv"
)

// GetEnvInt is used to load a given Environment variable
func GetEnvInt(key string, fallback int) int {
	rawValue, found := os.LookupEnv(key)
	if !found {
		return fallback
	}

	value, err := strconv.Atoi(rawValue)
	if err != nil {
		return fallback
	}

	return value
}
