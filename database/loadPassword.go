package database

import "github.com/Lol3rrr/sqlvault"

func (s *session) LoadPassword(ID string) (string, int64, error) {
	var result string
	var expiration int64
	err := s.SQLSession.WithRetry(func(con sqlvault.DB) error {
		loadQuery := `SELECT Password, Expiration FROM ` + s.PasswordsTable + `
		WHERE ID=$1;`
		err := con.QueryRow(loadQuery, ID).Scan(&result, &expiration)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return "", 0, err
	}

	return result, expiration, nil
}
