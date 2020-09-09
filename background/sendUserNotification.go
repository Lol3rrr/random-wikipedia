package background

import (
	"encoding/json"
	"random_wikipedia/general"
	"random_wikipedia/notifications"
	"random_wikipedia/wikipedia"
)

func sendUserNotification(user general.User, wSession wikipedia.Session, nSession notifications.Session) {
	ids := make([]int, 0, len(user.Lists))
	for _, tmpList := range user.Lists {
		ids = append(ids, tmpList.ID)
	}

	article, err := wSession.GetRandomArticle(ids)
	if err != nil {
		return
	}

	tmpNotification := general.ArticleNotification{
		ID:    article.ID,
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
