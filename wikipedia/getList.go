package wikipedia

func (s *session) getList(pageID int, title string) (list, error) {
	articles, continueKey, err := s.getArticlesInList(pageID, "")
	if err != nil {
		return list{}, err
	}

	for len(continueKey) > 0 {
		tmpArticles, tmpContinue, err := s.getArticlesInList(pageID, continueKey)
		if err != nil {
			return list{}, err
		}

		articles = append(articles, tmpArticles...)
		continueKey = tmpContinue
	}

	return list{
		ListInfo: ListInfo{
			ID:    pageID,
			Title: title,
		},
		Articles: articles,
	}, nil
}
