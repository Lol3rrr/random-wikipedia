package email

import (
	"errors"
	"net/smtp"

	"github.com/Lol3rrr/cvault"
)

// NewSession is used to obtain a new Session for emails
func NewSession(vault cvault.Session) (Session, error) {
	emailConfig, err := vault.ReadMap("/kv/data/wikipedia/email")
	if err != nil {
		return nil, err
	}

	server, worked := emailConfig["SMTPServer"].(string)
	if !worked {
		return nil, errors.New("Could not load 'SMTPServer' from Vault")
	}

	email, worked := emailConfig["Email"].(string)
	if !worked {
		return nil, errors.New("Could not load 'Email' from Vault")
	}

	password, worked := emailConfig["Password"].(string)
	if !worked {
		return nil, errors.New("Could not load 'Password' from Vault")
	}

	emailAuth := smtp.PlainAuth("", email, password, server)

	return &session{
		SMTPServer:  server,
		SourceEmail: email,
		Auth:        emailAuth,
	}, nil
}
