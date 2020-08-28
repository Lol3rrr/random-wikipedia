package email

import (
	"testing"

	"github.com/Lol3rrr/cvault"
	"github.com/Lol3rrr/cvault/mocks"
	"github.com/stretchr/testify/mock"
)

func TestNewSession(t *testing.T) {
	tables := []struct {
		Name       string
		InputVault cvault.Session
	}{
		{
			Name: "Valid",
			InputVault: &mocks.MockSession{
				Mock: mock.Mock{
					ExpectedCalls: []*mock.Call{
						{
							Method: "ReadMap",
							Arguments: mock.Arguments{
								"/kv/data/wikipedia/email",
							},
							ReturnArguments: mock.Arguments{
								map[string]interface{}{
									"SMTPServer": "testServer",
									"Email":      "testEmail",
									"Password":   "testPassword",
								},
								nil,
							},
						},
					},
				},
			},
		},
	}

	for _, table := range tables {
		inVault := table.InputVault

		t.Run(table.Name, func(t *testing.T) {
			NewSession(inVault)
		})
	}
}
