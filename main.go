package main

import (
	"random_wikipedia/wikipedia"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Starting...")

	wikiSession := wikipedia.NewSession()
	// Loading articles in unusual articles list
	articles, err := wikiSession.GetAllArticlesInList(154126)
	if err != nil {
		logrus.Errorf("Could not load articles: %s", err)
		return
	}

	for _, article := range articles {
		logrus.Infof("[%s] %s", article.Title, article.URL)
	}
	logrus.Infof("Loaded %d Articles", len(articles))
}
