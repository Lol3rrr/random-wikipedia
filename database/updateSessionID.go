package database

import "database/sql"

func (s *session) UpdateSessionID(ID, SessionID string) error {
	err := s.SQLSession.WithRetry(func(con *sql.DB) error {
		updateQuery := `UPDATE Users
		SET SessionID=$2
		WHERE ID=$1;`
		_, err := con.Exec(updateQuery, ID, SessionID)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
