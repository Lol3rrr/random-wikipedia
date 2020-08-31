package background

import (
	"encoding/json"
	"random_wikipedia/general"
	"random_wikipedia/notifications"
	"random_wikipedia/wikipedia"
)

func sendUserNotification(user general.User, wSession wikipedia.Session, nSession notifications.Session) {
	article, err := wSession.GetRandomArticle(user.Lists)
	if err != nil {
		return
	}

	tmpNotification := general.ArticleNotification{
		Title: article.Title,
		URL:   article.URL,
	}
	notifyBytes, err := json.Marshal(tmpNotification)
	if err != nil {
		return
	}

	err = nSession.SendNotification(notifyBytes, user)
	if err != nil {
		return
	}
}
