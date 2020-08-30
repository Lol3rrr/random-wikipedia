package database

import "database/sql"

func (s *session) DeletePassword(ID string) error {
	err := s.SQLSession.WithRetry(func(con *sql.DB) error {
		deleteQuery := `DELETE FROM ` +
			s.PasswordsTable +
			` WHERE ID=$1;`

		_, err := con.Exec(deleteQuery, ID)
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
