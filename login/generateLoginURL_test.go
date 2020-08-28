package login

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateLoginURL(t *testing.T) {
	tables := []struct {
		Name          string
		InputBaseURL  string
		InputEmail    string
		InputPassword string
		ResultURL     string
	}{
		{
			Name:          "Valid",
			InputBaseURL:  "https://example.com",
			InputEmail:    "test@example.com",
			InputPassword: "ADaasd132/1da)ADAS(DAsd!asd$",
			ResultURL:     "https://example.com/login/confirm?email=test%40example.com&password=ADaasd132%2F1da%29ADAS%28DAsd%21asd%24",
		},
		{
			Name:          "Valid with trailing / on baseURL",
			InputBaseURL:  "https://example.com/",
			InputEmail:    "test@example.com",
			InputPassword: "ADaasd132/1da)ADAS(DAsd!asd$",
			ResultURL:     "https://example.com/login/confirm?email=test%40example.com&password=ADaasd132%2F1da%29ADAS%28DAsd%21asd%24",
		},
	}

	for _, table := range tables {
		inBaseURL := table.InputBaseURL
		inEmail := table.InputEmail
		inPassword := table.InputPassword
		resURL := table.ResultURL

		t.Run(table.Name, func(t *testing.T) {
			outURL := generateLoginURL(inBaseURL, inEmail, inPassword)

			assert.Equal(t, resURL, outURL)
		})
	}
}
