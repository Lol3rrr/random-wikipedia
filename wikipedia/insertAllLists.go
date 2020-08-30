package wikipedia

import (
	"random_wikipedia/database"
)

func insertAllLists(dbSession database.Session, lists []ListInfo) error {
	for _, list := range lists {
		if err := dbSession.InsertList(list.ID, list.Title); err != nil {
			return err
		}
	}

	return nil
}
