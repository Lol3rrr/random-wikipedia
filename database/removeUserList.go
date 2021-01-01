package database

func (s *session) RemoveUserList(ID string, listID int) error {
	con, err := s.SQLSession.GetConnection()
	if err != nil {
		return err
	}

	deleteQuery := `DELETE
		FROM ` + s.UserlistsTable + `
		WHERE ID=$1 AND ListID=$2`

	_, err = con.Exec(deleteQuery, ID, listID)
	if err != nil {
		return err
	}

	return nil
}
