package database

import (
	"random_wikipedia/general"
)

func (s *session) InsertSettings(ID string, nSettings general.Settings, update bool) error {
	con, err := s.SQLSession.GetConnection()
	if err != nil {
		return err
	}

	insertQuery := `INSERT INTO ` + s.SettingsTable + ` (ID, NotifyTime)
		VALUES ($1, $2)
		ON CONFLICT (ID)
		DO NOTHING;`

	if update {
		insertQuery = `INSERT INTO ` + s.SettingsTable + ` (ID, NotifyTime)
			VALUES ($1, $2)
			ON CONFLICT (ID)
			DO UPDATE SET NotifyTime=$2;`
	}

	_, err = con.Exec(insertQuery, ID, nSettings.NotificationTime)
	if err != nil {
		return err
	}

	return nil

	return nil
}
