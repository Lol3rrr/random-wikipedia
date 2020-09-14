package api

import (
	"net/http"
	"random_wikipedia/general"

	"github.com/gofiber/fiber/v2"
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

func (a *api) handleLoadRandomArticle(ctx *fiber.Ctx) error {
	lists, err := loadLists(ctx)
	if err != nil {
		logrus.Errorf("[Article/Random] Loading Lists: %v", err)
		return ctx.SendStatus(http.StatusBadRequest)
	}

	article, err := a.WikipediaSession.GetRandomArticle(lists)
	if err != nil {
		logrus.Errorf("[Article/Random] Loading Random Article: %v", err)
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	articleResponse := general.ArticleNotification{
		ID:    article.ID,
		Title: article.Title,
		URL:   article.URL,
	}

	return ctx.JSON(articleResponse)
}
