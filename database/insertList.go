package database

import "database/sql"

func (s *session) InsertList(listID int, title string) error {
	err := s.SQLSession.WithRetry(func(con *sql.DB) error {
		insertQuery := `INSERT INTO ` + s.ListsTable + ` (ListID, Title)
		VALUES ($1, $2)
		ON CONFLICT (ListID)
		DO NOTHING;`

		_, err := con.Exec(insertQuery, listID, title)
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
