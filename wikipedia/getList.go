package wikipedia

func (s *session) GetList(pageID int, title string) (List, error) {
	articles, continueKey, err := s.GetArticlesInList(pageID, "")
	if err != nil {
		return &list{}, err
	}

	for len(continueKey) > 0 {
		tmpArticles, tmpContinue, err := s.GetArticlesInList(pageID, continueKey)
		if err != nil {
			return &list{}, err
		}

		articles = append(articles, tmpArticles...)
		continueKey = tmpContinue
	}

	return &list{
		ID:       pageID,
		Title:    title,
		Articles: articles,
	}, nil
}
