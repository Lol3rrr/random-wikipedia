package database

import "github.com/Lol3rrr/sqlvault"

func (s *session) RemoveUserFavorite(ID string, ArticleID int) error {
	err := s.SQLSession.WithRetry(func(con sqlvault.DB) error {
		deleteQuery := `DELETE
		FROM ` + s.FavoritesTable + `
		WHERE ID=$1 AND ArticleID=$2;`

		_, err := con.Exec(deleteQuery, ID, ArticleID)
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
