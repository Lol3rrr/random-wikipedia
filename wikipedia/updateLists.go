package wikipedia

import "github.com/sirupsen/logrus"

func (s *session) UpdateLists() {
	for key, list := range s.Lists {
		nList, err := s.getList(list.ID, list.Title)
		if err != nil {
			logrus.Errorf("Loading List: %v", err)
			continue
		}

		delete(s.Lists, key)
		s.Lists[key] = nList
	}
}
