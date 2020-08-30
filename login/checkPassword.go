package login

import (
	"errors"
	"time"
)

func (s *session) checkPassword(id, hashedPassword string) error {
	loadedPassword, expiration, err := s.DBSession.LoadPassword(id)
	if err != nil {
		return err
	}
	if time.Now().After(time.Unix(expiration, 0)) ||
		hashedPassword != loadedPassword {
		return errors.New("Invalid/Old Credentials")
	}

	return nil
}
