package database

func (s *session) InsertUserList(ID string, listID int) error {
	con, err := s.SQLSession.GetConnection()
	if err != nil {
		return err
	}

	insertQuery := `INSERT INTO ` + s.UserlistsTable + ` (ID, ListID)
		VALUES ($1, $2)
		ON CONFLICT (ID, ListID)
		DO NOTHING;`

	_, err = con.Query(insertQuery, ID, listID)
	if err != nil {
		return err
	}

	return nil
}
