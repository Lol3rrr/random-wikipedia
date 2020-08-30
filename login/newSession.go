package login

import (
	"random_wikipedia/database"
	"random_wikipedia/login/email"

	"github.com/Lol3rrr/cvault"
)

// NewSession is used to obtain a new Session for Login stuff
func NewSession(baseURL string, vault cvault.Session, dbSession database.Session) (Session, error) {
	emailSession, err := email.NewSession(vault)
	if err != nil {
		return nil, err
	}

	return &session{
		VaultSession: vault,
		DBSession:    dbSession,
		Email:        emailSession,
		BaseURL:      baseURL,
	}, nil
}
