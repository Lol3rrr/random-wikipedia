package database

import (
	"random_wikipedia/general"
)

func (s *session) LoadUserSessionID(SessionID string) (general.User, error) {
	con, err := s.SQLSession.GetConnection()
	if err != nil {
		return general.User{}, err
	}

	var id string
	loadIDQuery := `SELECT ID FROM ` + s.UsersTable + ` WHERE SessionID=$1;`
	err = con.QueryRow(loadIDQuery, SessionID).Scan(&id)
	if err != nil {
		return general.User{}, err
	}

	return s.LoadUserID(id)
}
