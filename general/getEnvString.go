package general

import "os"

// GetEnvString is used to load a given Environment variable
func GetEnvString(key, fallback string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return fallback
	}

	return value
}
