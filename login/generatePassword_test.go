package login

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePassword(t *testing.T) {
	tables := []struct {
		Name        string
		InputLength int
	}{
		{
			Name:        "Valid",
			InputLength: 20,
		},
	}

	for _, table := range tables {
		inLength := table.InputLength

		t.Run(table.Name, func(t *testing.T) {
			outPassword := generatePassword(inLength)

			assert.Equal(t, inLength, len(outPassword))
		})
	}
}
