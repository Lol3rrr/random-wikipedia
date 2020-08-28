package login

import (
	"random_wikipedia/general"

	"github.com/Lol3rrr/cvault"
)

// Session represents a simple abstraction for all the login stuff
type Session interface {
	Login(email, password string) bool
	LoadUser(sessionID string) (general.User, error)
}

type session struct {
	VaultSession cvault.Session
}
