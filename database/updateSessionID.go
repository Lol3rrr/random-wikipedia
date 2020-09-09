package database

import "github.com/Lol3rrr/sqlvault"

func (s *session) UpdateSessionID(ID, SessionID string) error {
	err := s.SQLSession.WithRetry(func(con sqlvault.DB) error {
		updateQuery := `UPDATE ` + s.UsersTable + `
		SET SessionID=$2
		WHERE ID=$1;`
		_, err := con.Exec(updateQuery, ID, SessionID)
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
