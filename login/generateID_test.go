package login

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	tables := []struct {
		Name       string
		InputEmail string
		Result     string
	}{
		{
			Name:       "Valid",
			InputEmail: "test@example.com",
			Result:     "lz3-Rj7IV4X1-Vr1ujkG7tstkxwk5pgkqJ6mXbpOgTs=",
		},
		{
			Name:       "Valid, other email",
			InputEmail: "test@example.net",
			Result:     "qv41gg4-5GuUPzOX_z_EYqc7jSDKQCkYvSHhBymkZoM=",
		},
	}

	for _, table := range tables {
		inEmail := table.InputEmail
		res := table.Result

		t.Run(table.Name, func(t *testing.T) {
			out := generateID(inEmail)

			assert.Equal(t, res, out)
		})
	}
}
