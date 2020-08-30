package general

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvInt(t *testing.T) {
	tables := []struct {
		Name          string
		EnvKey        string
		EnvValue      string
		InputKey      string
		InputFallback int
		Result        int
	}{
		{
			Name:          "Valid",
			EnvKey:        "testKey",
			EnvValue:      "1",
			InputKey:      "testKey",
			InputFallback: -1,
			Result:        1,
		},
		{
			Name:          "Key is missing",
			EnvKey:        "otherKey",
			EnvValue:      "1",
			InputKey:      "testKey",
			InputFallback: -1,
			Result:        -1,
		},
		{
			Name:          "Value is not an int",
			EnvKey:        "testKey",
			EnvValue:      "test",
			InputKey:      "testKey",
			InputFallback: -1,
			Result:        -1,
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
			out := GetEnvInt(inKey, inFallback)

			assert.Equal(t, res, out)
		})
	}
}
