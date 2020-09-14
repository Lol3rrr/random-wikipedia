package api

import (
	"net/http"
	"random_wikipedia/general"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (a *api) handleDeleteFavorite(ctx *fiber.Ctx) error {
	var query map[string]int
	err := ctx.BodyParser(&query)
	if err != nil {
		logrus.Errorf("[Settings/Favorites/Delete] Parsing Body: %v", err)
		return ctx.SendStatus(http.StatusBadRequest)
	}

	articleID, found := query["articleID"]
	if !found {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	user := ctx.Locals("user").(general.User)

	err = a.DBSession.RemoveUserFavorite(user.ID, articleID)
	if err != nil {
		logrus.Errorf("[Settings/Favorites/Delete] Removing User-Favorite: %v", err)
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.SendStatus(http.StatusOK)
}
