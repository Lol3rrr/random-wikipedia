package database

func (s *session) InsertUserFavorite(ID string, ArticleID int, ArticleTitle, ArticleURL string) error {
	con, err := s.SQLSession.GetConnection()
	if err != nil {
		return err
	}

	insertFavArticleQuery := `INSERT INTO ` +
		s.FavArticlesTable + ` (ArticleID, Title, URL)
			VALUES ($1, $2, $3)
			ON CONFLICT
			DO NOTHING;`

	_, err = con.Exec(insertFavArticleQuery, ArticleID, ArticleTitle, ArticleURL)
	if err != nil {
		return err
	}

	insertFavoriteQuery := `INSERT INTO ` +
		s.FavoritesTable + ` (ID, ArticleID)
			VALUES ($1, $2)
			ON CONFLICT
			DO NOTHING;`

	_, err = con.Exec(insertFavoriteQuery, ID, ArticleID)
	if err != nil {
		return err
	}

	return nil
}
