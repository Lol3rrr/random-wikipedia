package database

func (s *session) InsertSubscription(ID, subscription string, update bool) error {
	con, err := s.SQLSession.GetConnection()
	if err != nil {
		return err
	}

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

	_, err = con.Exec(insertQuery, ID, subscription)
	if err != nil {
		return err
	}

	return nil
}
