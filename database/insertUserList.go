package database

import "database/sql"

func (s *session) InsertUserList(ID string, listID int) error {
	err := s.SQLSession.WithRetry(func(con *sql.DB) error {
		insertQuery := `INSERT INTO ` + s.UserlistsTable + ` (ID, ListID)
		VALUES ($1, $2)
		ON CONFLICT (ID, ListID)
		DO NOTHING;`

		_, err := con.Query(insertQuery, ID, listID)
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
