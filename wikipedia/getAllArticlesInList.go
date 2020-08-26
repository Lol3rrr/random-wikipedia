package wikipedia

func (s *session) GetAllArticlesInList(pageID int) ([]Article, error) {
	articles, continueKey, err := s.GetArticlesInList(pageID, "")
	if err != nil {
		return nil, err
	}

	for len(continueKey) > 0 {
		tmpArticles, tmpContinue, err := s.GetArticlesInList(pageID, continueKey)
		if err != nil {
			return nil, err
		}

		articles = append(articles, tmpArticles...)
		continueKey = tmpContinue
	}

	return articles, nil
}
