package database

import "github.com/Lol3rrr/sqlvault"

func (s *session) InsertUser(ID, Email string) error {
	err := s.SQLSession.WithRetry(func(con sqlvault.DB) error {
		insertQuery := `INSERT INTO ` + s.UsersTable + ` (ID, Email, SessionID)
		VALUES ($1, $2, $3)
		ON CONFLICT
		DO NOTHING;`

		_, err := con.Exec(insertQuery, ID, Email, "")
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
