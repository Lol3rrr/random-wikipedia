package database

import (
	"database/sql"
	"random_wikipedia/general"
)

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

		loadListsQuery := `SELECT ListID FROM ` + s.UserlistsTable + ` WHERE ID=$1;`
		rows, err := con.Query(loadListsQuery, ID)
		if err != nil {
			return err
		}

		for rows.Next() {
			var tmpListID int
			if err := rows.Scan(&tmpListID); err != nil {
				return err
			}

			result.Lists = append(result.Lists, tmpListID)
		}

		return nil
	})
	if err != nil {
		return general.User{}, err
	}

	return result, nil
}
