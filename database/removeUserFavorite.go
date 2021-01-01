package database

func (s *session) RemoveUserFavorite(ID string, ArticleID int) error {
	con, err := s.SQLSession.GetConnection()
	if err != nil {
		return err
	}

	deleteQuery := `DELETE
		FROM ` + s.FavoritesTable + `
		WHERE ID=$1 AND ArticleID=$2;`

	_, err = con.Exec(deleteQuery, ID, ArticleID)
	if err != nil {
		return err
	}

	return nil
}
