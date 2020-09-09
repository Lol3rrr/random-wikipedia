package database

import "github.com/Lol3rrr/sqlvault"

func (s *session) RemoveUserList(ID string, listID int) error {
	err := s.SQLSession.WithRetry(func(con sqlvault.DB) error {
		deleteQuery := `DELETE
		FROM ` + s.UserlistsTable + `
		WHERE ID=$1 AND ListID=$2`

		_, err := con.Exec(deleteQuery, ID, listID)
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
