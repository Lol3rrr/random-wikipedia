package wikipedia

import (
	"errors"
	"math/rand"
)

func (s *session) GetRandomArticle(listIDs []int) (Article, error) {
	listSize := len(listIDs)
	if listSize <= 0 {
		return Article{}, errors.New("No lists possible")
	}
	listID := listIDs[rand.Intn(listSize)]

	list, found := s.Lists[listID]
	if !found {
		return Article{}, errors.New("Could not find a list")
	}

	return list.GetRandomArticle()
}
