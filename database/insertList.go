package database

func (s *session) InsertList(listID int, title string) error {
	con, err := s.SQLSession.GetConnection()
	if err != nil {
		return err
	}

	insertQuery := `INSERT INTO ` + s.ListsTable + ` (ListID, Title)
		VALUES ($1, $2)
		ON CONFLICT (ListID)
		DO NOTHING;`

	_, err = con.Exec(insertQuery, listID, title)
	if err != nil {
		return err
	}

	return nil
}
