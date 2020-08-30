package login

import (
	"errors"
	"random_wikipedia/database"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheckPassword(t *testing.T) {
	tables := []struct {
		Name                string
		InputSession        *session
		InputID             string
		InputHashedPassword string
		ResultError         bool
	}{
		{
			Name: "Valid",
			InputSession: &session{
				DBSession: &database.MockDatabase{
					Mock: mock.Mock{
						ExpectedCalls: []*mock.Call{
							{
								Method: "LoadPassword",
								Arguments: mock.Arguments{
									"testID",
								},
								ReturnArguments: mock.Arguments{
									"testPassword",
									time.Now().Add(10 * time.Minute).Unix(),
									nil,
								},
							},
						},
					},
				},
			},
			InputID:             "testID",
			InputHashedPassword: "testPassword",
			ResultError:         false,
		},
		{
			Name: "Database returns error",
			InputSession: &session{
				DBSession: &database.MockDatabase{
					Mock: mock.Mock{
						ExpectedCalls: []*mock.Call{
							{
								Method: "LoadPassword",
								Arguments: mock.Arguments{
									"testID",
								},
								ReturnArguments: mock.Arguments{
									"",
									int64(0),
									errors.New("testError"),
								},
							},
						},
					},
				},
			},
			InputID:             "testID",
			InputHashedPassword: "testPassword",
			ResultError:         true,
		},
		{
			Name: "The returned password dont match",
			InputSession: &session{
				DBSession: &database.MockDatabase{
					Mock: mock.Mock{
						ExpectedCalls: []*mock.Call{
							{
								Method: "LoadPassword",
								Arguments: mock.Arguments{
									"testID",
								},
								ReturnArguments: mock.Arguments{
									"dbPassword",
									time.Now().Add(10 * time.Minute).Unix(),
									nil,
								},
							},
						},
					},
				},
			},
			InputID:             "testID",
			InputHashedPassword: "testPassword",
			ResultError:         true,
		},
		{
			Name: "Password has expired",
			InputSession: &session{
				DBSession: &database.MockDatabase{
					Mock: mock.Mock{
						ExpectedCalls: []*mock.Call{
							{
								Method: "LoadPassword",
								Arguments: mock.Arguments{
									"testID",
								},
								ReturnArguments: mock.Arguments{
									"testPassword",
									time.Now().Add(-10 * time.Minute).Unix(),
									nil,
								},
							},
						},
					},
				},
			},
			InputID:             "testID",
			InputHashedPassword: "testPassword",
			ResultError:         true,
		},
	}

	for _, table := range tables {
		inSession := table.InputSession
		inID := table.InputID
		inHashedPassword := table.InputHashedPassword
		resError := table.ResultError

		t.Run(table.Name, func(t *testing.T) {
			outErr := inSession.checkPassword(inID, inHashedPassword)

			if resError {
				assert.NotNil(t, outErr)
			} else {
				assert.Nil(t, outErr)
			}
		})
	}
}
