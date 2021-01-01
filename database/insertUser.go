package database

func (s *session) InsertUser(ID, Email string) error {
	con, err := s.SQLSession.GetConnection()
	if err != nil {
		return err
	}

	insertQuery := `INSERT INTO ` + s.UsersTable + ` (ID, Email, SessionID)
		VALUES ($1, $2, $3)
		ON CONFLICT
		DO NOTHING;`

	_, err = con.Exec(insertQuery, ID, Email, "")
	if err != nil {
		return err
	}

	return nil
}
