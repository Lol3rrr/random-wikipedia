package general

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvString(t *testing.T) {
	tables := []struct {
		Name          string
		EnvKey        string
		EnvValue      string
		InputKey      string
		InputFallback string
		Result        string
	}{
		{
			Name:          "Valid",
			EnvKey:        "testKey",
			EnvValue:      "testValue",
			InputKey:      "testKey",
			InputFallback: "testFallback",
			Result:        "testValue",
		},
		{
			Name:          "Key is missing",
			EnvKey:        "otherKey",
			EnvValue:      "testValue",
			InputKey:      "testKey",
			InputFallback: "testFallback",
			Result:        "testFallback",
		},
	}

	for _, table := range tables {
		envKey := table.EnvKey
		envValue := table.EnvValue
		inKey := table.InputKey
		inFallback := table.InputFallback
		res := table.Result

		t.Run(table.Name, func(t *testing.T) {
			os.Clearenv()
			os.Setenv(envKey, envValue)
			out := GetEnvString(inKey, inFallback)

			assert.Equal(t, res, out)
		})
	}
}
