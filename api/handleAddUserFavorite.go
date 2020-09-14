package api

import (
	"net/http"
	"random_wikipedia/general"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (a *api) handleAddUserFavorite(ctx *fiber.Ctx) error {
	var articleQuery general.Article
	err := ctx.BodyParser(&articleQuery)
	if err != nil {
		logrus.Errorf("[Settings/Favorites/Add] Parsing Body: %v", err)
		return ctx.SendStatus(http.StatusBadRequest)
	}

	user := ctx.Locals("user").(general.User)

	err = a.DBSession.InsertUserFavorite(user.ID, articleQuery.ID, articleQuery.Title, articleQuery.URL)
	if err != nil {
		logrus.Errorf("[Settings/Favorites/Add] Inserting Favorite: %v", err)
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.SendStatus(http.StatusOK)
}
