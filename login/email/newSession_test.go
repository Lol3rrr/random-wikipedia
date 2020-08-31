package email

import (
	"errors"
	"testing"

	"github.com/Lol3rrr/cvault"
	"github.com/Lol3rrr/cvault/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewSession(t *testing.T) {
	tables := []struct {
		Name        string
		InputVault  cvault.Session
		ResultError bool
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
									"ServerPort": "25",
									"Email":      "testEmail",
									"Password":   "testPassword",
								},
								nil,
							},
						},
					},
				},
			},
			ResultError: false,
		},
		{
			Name: "Vault returns error",
			InputVault: &mocks.MockSession{
				Mock: mock.Mock{
					ExpectedCalls: []*mock.Call{
						{
							Method: "ReadMap",
							Arguments: mock.Arguments{
								"/kv/data/wikipedia/email",
							},
							ReturnArguments: mock.Arguments{
								map[string]interface{}{},
								errors.New("testError"),
							},
						},
					},
				},
			},
			ResultError: true,
		},
		{
			Name: "Missing password",
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
									"ServerPort": "25",
									"Email":      "testEmail",
								},
								nil,
							},
						},
					},
				},
			},
			ResultError: true,
		},
		{
			Name: "Missing Email",
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
									"ServerPort": "25",
									"Password":   "testPassword",
								},
								nil,
							},
						},
					},
				},
			},
			ResultError: true,
		},
		{
			Name: "Missing ServerPort",
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
			ResultError: true,
		},
		{
			Name: "Missing SMTPServer",
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
									"ServerPort": "25",
									"Email":      "testEmail",
									"Password":   "testPassword",
								},
								nil,
							},
						},
					},
				},
			},
			ResultError: true,
		},
	}

	for _, table := range tables {
		inVault := table.InputVault
		resErr := table.ResultError

		t.Run(table.Name, func(t *testing.T) {
			_, outError := NewSession(inVault)

			if resErr {
				assert.NotNil(t, outError)
			} else {
				assert.Nil(t, outError)
			}
		})
	}
}
