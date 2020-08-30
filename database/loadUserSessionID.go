package database

import (
	"database/sql"
	"random_wikipedia/general"
)

func (s *session) LoadUserSessionID(SessionID string) (general.User, error) {
	var id string
	err := s.SQLSession.WithRetry(func(con *sql.DB) error {
		loadIDQuery := `SELECT ID FROM ` + s.UsersTable + ` WHERE SessionID=$1;`
		err := con.QueryRow(loadIDQuery, SessionID).Scan(&id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return general.User{}, err
	}

	return s.LoadUserID(id)
}
