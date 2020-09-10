package api

import (
	"net/http"
	"random_wikipedia/general"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (a *api) handleDeleteFavorite(ctx *fiber.Ctx) {
	var query map[string]int
	err := ctx.BodyParser(&query)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		logrus.Errorf("[Settings/Favorites/Delete] Parsing Body: %v", err)
		return
	}

	articleID, found := query["articleID"]
	if !found {
		ctx.SendStatus(http.StatusBadRequest)
		return
	}

	user := ctx.Locals("user").(general.User)

	err = a.DBSession.RemoveUserFavorite(user.ID, articleID)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		logrus.Errorf("[Settings/Favorites/Delete] Removing User-Favorite: %v", err)
		return
	}

	ctx.SendStatus(http.StatusOK)
}
