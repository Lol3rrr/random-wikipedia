package database

import "github.com/Lol3rrr/sqlvault"

func (s *session) InsertUserFavorite(ID string, ArticleID int, ArticleTitle, ArticleURL string) error {
	err := s.SQLSession.WithRetry(func(con sqlvault.DB) error {
		insertFavArticleQuery := `INSERT INTO ` +
			s.FavArticlesTable + ` (ArticleID, Title, URL)
			VALUES ($1, $2, $3)
			ON CONFLICT
			DO NOTHING;`

		_, err := con.Exec(insertFavArticleQuery, ArticleID, ArticleTitle, ArticleURL)
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
	})
	if err != nil {
		return err
	}

	return nil
}
