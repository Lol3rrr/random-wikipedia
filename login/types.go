package login

import (
	"random_wikipedia/database"
	"random_wikipedia/login/email"

	"github.com/Lol3rrr/cvault"
)

// Session represents a simple abstraction for all the login stuff
type Session interface {
	Login(email string) error
	// Authenticate returns the new sessionID for the user
	Authenticate(email, password string) (string, error)
}

type session struct {
	VaultSession cvault.Session
	DBSession    database.Session
	Email        email.Session
	BaseURL      string
}
