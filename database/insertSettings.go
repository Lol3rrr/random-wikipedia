package database

import (
	"random_wikipedia/general"

	"github.com/Lol3rrr/sqlvault"
)

func (s *session) InsertSettings(ID string, nSettings general.Settings, update bool) error {
	err := s.SQLSession.WithRetry(func(con sqlvault.DB) error {
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
