package email

import "net/smtp"

// Session is a simple abstraction for dealing with emails
type Session interface {
	// SendLoginEmail is used to send a simple Login email to the
	// given email address with the LoginURL in it
	SendLoginEmail(receiver, loginURL string) error
}

type session struct {
	SMTPServer  string
	Auth        smtp.Auth
	SourceEmail string
}
