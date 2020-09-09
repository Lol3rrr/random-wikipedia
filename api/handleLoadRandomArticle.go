package api

import (
	"net/http"
	"random_wikipedia/general"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func loadLists(ctx *fiber.Ctx) ([]int, error) {
	var query map[string][]int
	err := ctx.BodyParser(&query)
	if err != nil {
		return nil, err
	}

	return query["lists"], nil
}

func (a *api) handleLoadRandomArticle(ctx *fiber.Ctx) {
	lists, err := loadLists(ctx)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		logrus.Errorf("[Article/Random] Loading Lists: %v", err)
		return
	}

	article, err := a.WikipediaSession.GetRandomArticle(lists)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		logrus.Errorf("[Article/Random] Loading Random Article: %v", err)
		return
	}

	articleResponse := general.ArticleNotification{
		ID:    article.ID,
		Title: article.Title,
		URL:   article.URL,
	}

	err = ctx.JSON(articleResponse)
	if err != nil {
		logrus.Errorf("[Article/Random] Sending JSON: %v", err)
	}
}
