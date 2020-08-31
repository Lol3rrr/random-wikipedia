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
	if !worked || len(server) <= 0 {
		return nil, errors.New("Could not load 'SMTPServer' from Vault")
	}

	port, worked := emailConfig["ServerPort"].(string)
	if !worked || len(port) <= 0 {
		return nil, errors.New("Could not load 'ServerPort' from Vault")
	}

	email, worked := emailConfig["Email"].(string)
	if !worked || len(email) <= 0 {
		return nil, errors.New("Could not load 'Email' from Vault")
	}

	password, worked := emailConfig["Password"].(string)
	if !worked || len(password) <= 0 {
		return nil, errors.New("Could not load 'Password' from Vault")
	}

	emailAuth := smtp.PlainAuth("", email, password, server)

	return &session{
		SMTPServer:  server + ":" + port,
		SourceEmail: email,
		Auth:        emailAuth,
	}, nil
}
