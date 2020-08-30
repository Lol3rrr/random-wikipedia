package database

import (
	"database/sql"
	"random_wikipedia/general"
)

func (s *session) InsertSettings(ID string, nSettings general.Settings, update bool) error {
	err := s.SQLSession.WithRetry(func(con *sql.DB) error {
		insertQuery := `INSERT INTO ` + s.SettingsTable + ` (ID, NotifyTime)
		VALUES ($1, $2)
		ON CONFLICT (ID)
		DO NOTHING;`

		if update {
			insertQuery = `INSERT INTO ` + s.SettingsTable + ` (ID, NotifyTime)
			VALUES ($1, $2)
			ON CONFLICT (ID)
			DO SET NotifyTime=$2;`
		}

		_, err := con.Exec(insertQuery, ID, nSettings.NotificationTime)
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
