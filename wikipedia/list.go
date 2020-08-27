package wikipedia

import (
	"errors"
	"math/rand"
	"time"
)

func (l *list) GetID() int {
	return l.ID
}
func (l *list) GetTitle() string {
	return l.Title
}
func (l *list) GetArticles() []Article {
	return l.Articles
}

func (l *list) GetRandomArticle() (Article, error) {
	length := len(l.Articles)
	if length <= 0 {
		return Article{}, errors.New("List does not contain any Articles")
	}

	rand.Seed(time.Now().Unix())
	return l.Articles[rand.Intn(length)], nil
}
