package database

func (s *session) DeletePassword(ID string) error {
	con, err := s.SQLSession.GetConnection()
	if err != nil {
		return err
	}

	deleteQuery := `DELETE FROM ` +
		s.PasswordsTable +
		` WHERE ID=$1;`

	_, err = con.Exec(deleteQuery, ID)
	if err != nil {
		return err
	}

	return nil
}
