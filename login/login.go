package login

import (
	"time"
)

func (s *session) Login(email string) error {
	password := generatePassword(40)
	url := generateLoginURL(s.BaseURL, email, password)

	id := generateID(email)
	hashedPassword := hashPassword(password)

	passwordExpiration := time.Now().Add(30 * time.Minute).Unix()

	err := s.DBSession.InsertUser(id, email)
	if err != nil {
		return err
	}
	err = s.DBSession.InsertPassword(id, hashedPassword, passwordExpiration)
	if err != nil {
		return err
	}

	return s.Email.SendLoginEmail(email, url)
}
