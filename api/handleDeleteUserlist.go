package api

import (
	"net/http"
	"random_wikipedia/general"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (a *api) handleDeleteUserlist(ctx *fiber.Ctx) error {
	var query map[string]int
	err := ctx.BodyParser(&query)
	if err != nil {
		logrus.Errorf("[Settings/Lists/Delete] Parsing Body: %v", err)
		return ctx.SendStatus(http.StatusBadRequest)
	}

	listID, found := query["listID"]
	if !found {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	user := ctx.Locals("user").(general.User)
	err = a.DBSession.RemoveUserList(user.ID, listID)
	if err != nil {
		logrus.Errorf("[Settings/Lists/Delete] Removing User-List: %v", err)
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.SendStatus(http.StatusOK)
}
