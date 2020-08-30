package wikipedia

import (
	"random_wikipedia/database"

	"github.com/sirupsen/logrus"
)

// NewSession is used to obtain a new Session
func NewSession(lists []ListInfo, dbSession database.Session) Session {
	tmp := &session{
		BaseURL:   "https://en.wikipedia.org/",
		UserAgent: "Random-Wikipedia/0.1 (;leon@lol3r.net)",
		Lists:     make(map[int]list),
	}

	for _, tmpList := range lists {
		resList, err := tmp.getList(tmpList.ID, tmpList.Title)
		if err != nil {
			logrus.Errorf("Loading List: %v", err)
			continue
		}

		tmp.Lists[tmpList.ID] = resList
	}

	err := insertAllLists(dbSession, lists)
	if err != nil {
		logrus.Errorf("Inserting Lists: %v", err)
	}

	return tmp
}
