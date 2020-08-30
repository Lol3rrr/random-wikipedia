package login

import (
	"random_wikipedia/general"
)

func (s *session) Authenticate(email, password string) (string, error) {
	id := generateID(email)
	hashedPassword := hashPassword(password)
	defer s.DBSession.DeletePassword(id)

	err := s.checkPassword(id, hashedPassword)
	if err != nil {
		return "", err
	}

	sessionID := generateSessionID()
	err = s.DBSession.UpdateSessionID(id, sessionID)
	if err != nil {
		return "", err
	}

	err = s.DBSession.InsertSubscription(id, "", false)
	if err != nil {
		return "", err
	}

	err = s.DBSession.InsertSettings(id, general.Settings{
		NotificationTime: -1,
	}, false)
	if err != nil {
		return "", err
	}

	return sessionID, nil
}
