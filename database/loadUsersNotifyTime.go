package database

import (
	"random_wikipedia/general"
)

func (s *session) LoadUsersNotifyTime(notifyTime int) ([]general.User, error) {
	result := make([]general.User, 0)

	con, err := s.SQLSession.GetConnection()
	if err != nil {
		return nil, err
	}

	loadQuery := `SELECT S.ID, N.Subscription, S.NotifyTime
		FROM ` + s.SettingsTable + ` AS S
		INNER JOIN ` + s.NotificationsTable + ` AS N
		ON S.ID=N.ID
		WHERE S.NotifyTime=$1`
	rows, err := con.Query(loadQuery, notifyTime)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id string
		var subscription string
		var notifyTime int = -1
		if err := rows.Scan(&id, &subscription, &notifyTime); err != nil {
			continue
		}
		if len(id) <= 0 || len(subscription) <= 0 || notifyTime < 0 {
			continue
		}

		lists, err := queryUserLists(con, s.UserlistsTable, s.ListsTable, id)
		if err != nil {
			continue
		}

		result = append(result, general.User{
			ID:           id,
			Subscription: subscription,
			Settings: general.Settings{
				NotificationTime: notifyTime,
			},
			Lists: lists,
		})
	}

	return result, nil
}
