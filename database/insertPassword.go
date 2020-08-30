package database

import "database/sql"

func (s *session) InsertPassword(ID, Password string, Expiration int64) error {
	err := s.SQLSession.WithRetry(func(con *sql.DB) error {
		insertQuery := `INSERT INTO ` + s.PasswordsTable + ` (ID, Password, Expiration)
		VALUES ($1, $2, $3)
		ON CONFLICT (ID)
		DO UPDATE SET Password=$2, Expiration=$3;`

		_, err := con.Exec(insertQuery, ID, Password, Expiration)
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
