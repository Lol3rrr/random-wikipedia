package database

import (
	"random_wikipedia/general"

	"github.com/Lol3rrr/sqlvault"
)

func queryUserLists(con sqlvault.DB, userListsTable, listsTable, userID string) ([]general.List, error) {
	result := make([]general.List, 0)

	loadListsQuery := `SELECT L.ListID, L.Title
	FROM ` + userListsTable + ` as UL
	INNER JOIN ` + listsTable + ` as L
	ON UL.ListID=L.ListID
	WHERE UL.ID=$1;`
	rows, err := con.Query(loadListsQuery, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmpList general.List
		if err := rows.Scan(&tmpList.ID, &tmpList.Name); err != nil {
			return nil, err
		}
		result = append(result, tmpList)
	}

	return result, nil
}
