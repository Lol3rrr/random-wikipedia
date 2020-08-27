package main

import (
	"random_wikipedia/wikipedia"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Starting...")

	wikiSession := wikipedia.NewSession()
	// Loading articles in unusual articles list
	unusualArticles, err := wikiSession.GetList(154126, "Unusual Articles")
	if err != nil {
		logrus.Errorf("Could not load articles: %s", err)
		return
	}

	for _, article := range unusualArticles.GetArticles() {
		logrus.Infof("[%s] %s", article.Title, article.URL)
	}
	logrus.Infof("Loaded %d Articles", len(unusualArticles.GetArticles()))

	randomArticle, _ := unusualArticles.GetRandomArticle()
	logrus.Infof("Random Article: '%+v'", randomArticle)
}
