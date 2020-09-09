package database

import (
	"database/sql"
	"random_wikipedia/general"
)

func loadUserFavorites(con *sql.DB, favoriteTable, favArticlesTable, userID string) ([]general.Article, error) {
	result := make([]general.Article, 0)

	loadFavoritesQuery := `SELECT FA.ArticleID, FA.Title, FA.URL
		FROM ` + favArticlesTable + ` as FA
		INNER JOIN ` + favoriteTable + ` as F
		ON F.ArticleID=FA.ArticleID
		WHERE F.ID=$1;`

	rows, err := con.Query(loadFavoritesQuery, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmpFavorite general.Article
		if err := rows.Scan(&tmpFavorite.ID, &tmpFavorite.Title, &tmpFavorite.URL); err != nil {
			return nil, err
		}
		result = append(result, tmpFavorite)
	}

	return result, nil
}

func (s *session) LoadUserID(ID string) (general.User, error) {
	result := general.User{
		ID: ID,
	}

	err := s.SQLSession.WithRetry(func(con *sql.DB) error {
		loadSubscriptionQuery := `SELECT N.Subscription, S.NotifyTime
		FROM ` + s.NotificationsTable + ` as N
		INNER JOIN ` + s.SettingsTable + ` as S
		ON N.ID=S.ID
		WHERE N.ID=$1;`
		err := con.QueryRow(loadSubscriptionQuery, ID).Scan(&result.Subscription, &result.Settings.NotificationTime)
		if err != nil {
			return err
		}

		userLists, err := queryUserLists(con, s.UserlistsTable, s.ListsTable, ID)
		if err != nil {
			return err
		}
		result.Lists = userLists

		userFavorites, err := loadUserFavorites(con, s.FavoritesTable, s.FavArticlesTable, ID)
		if err != nil {
			return err
		}
		result.Favorites = userFavorites

		return nil
	})
	if err != nil {
		return general.User{}, err
	}

	return result, nil
}
