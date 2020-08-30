package email

import "net/smtp"

func (s *session) SendLoginEmail(receiver, loginURL string) error {

	msg := []byte("Subject: One Time Login \r\n" +
		"\r\n" +
		"Click on the following Link to automatically log in: " + loginURL + "\r\n")

	return smtp.SendMail(s.SMTPServer, s.Auth, s.SourceEmail, []string{receiver}, msg)
}
