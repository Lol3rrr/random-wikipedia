package database

func (s *session) UpdateSessionID(ID, SessionID string) error {
	con, err := s.SQLSession.GetConnection()
	if err != nil {
		return err
	}

	updateQuery := `UPDATE ` + s.UsersTable + `
		SET SessionID=$2
		WHERE ID=$1;`
	_, err = con.Exec(updateQuery, ID, SessionID)
	if err != nil {
		return err
	}

	return nil
}
