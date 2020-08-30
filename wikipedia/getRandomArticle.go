package wikipedia

import (
	"errors"
	"math/rand"
)

func (s *session) GetRandomArticle(listIDs []int) (Article, error) {
	listID := listIDs[rand.Intn(len(listIDs))]

	list, found := s.Lists[listID]
	if !found {
		return Article{}, errors.New("Could not find a list")
	}

	return list.GetRandomArticle()
}
