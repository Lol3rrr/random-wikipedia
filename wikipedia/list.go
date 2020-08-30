package wikipedia

import (
	"errors"
	"math/rand"
	"time"
)

func (l *list) GetRandomArticle() (Article, error) {
	length := len(l.Articles)
	if length <= 0 {
		return Article{}, errors.New("List does not contain any Articles")
	}

	rand.Seed(time.Now().Unix())
	return l.Articles[rand.Intn(length)], nil
}
