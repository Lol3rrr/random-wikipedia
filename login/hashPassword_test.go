package login

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	tables := []struct {
		Name          string
		InputPassword string
		Result        string
	}{
		{
			Name:          "Valid",
			InputPassword: "testPassword",
			Result:        "iluLRhHe5Gs9rzUx-rsqc6k6K-N26qJA3BFd1YGL0kpTPu7ppGqqJ8gGRRbkieYLdVM1Bud04ZeSKEKMkQrydQ==",
		},
		{
			Name:          "Valid, other Password",
			InputPassword: "otherPassword",
			Result:        "6S2kG_b6oHgWBXQjKDKTayXWu2cs9374lxFrL9uVpmYUlq0lw_ZFU9svMtYVDV5aVjJqRbLWZ_dfeaaJwYxkhQ==",
		},
	}

	for _, table := range tables {
		inPassword := table.InputPassword
		res := table.Result

		t.Run(table.Name, func(t *testing.T) {
			out := hashPassword(inPassword)

			assert.Equal(t, res, out)
		})
	}
}
