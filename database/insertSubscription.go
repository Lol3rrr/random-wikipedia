package database

import "database/sql"

func (s *session) InsertSubscription(ID, subscription string, update bool) error {
	err := s.SQLSession.WithRetry(func(con *sql.DB) error {
		insertQuery := `INSERT INTO ` + s.NotificationsTable + ` (ID, Subscription)
		VALUES ($1, $2)
		ON CONFLICT (ID)
		DO NOTHING;`

		if update {
			insertQuery = `INSERT INTO ` + s.NotificationsTable + ` (ID, Subscription)
			VALUES ($1, $2)
			ON CONFLICT (ID)
			DO UPDATE SET Subscription=$2;`
		}

		_, err := con.Exec(insertQuery, ID, subscription)
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
