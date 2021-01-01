package database

func (s *session) LoadPassword(ID string) (string, int64, error) {
	con, err := s.SQLSession.GetConnection()
	if err != nil {
		return "", 0, err
	}

	var result string
	var expiration int64
	loadQuery := `SELECT Password, Expiration FROM ` + s.PasswordsTable + `
		WHERE ID=$1;`
	err = con.QueryRow(loadQuery, ID).Scan(&result, &expiration)
	if err != nil {
		return "", 0, err
	}

	return result, expiration, nil
}
