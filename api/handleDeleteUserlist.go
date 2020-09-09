package api

import (
	"net/http"
	"random_wikipedia/general"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (a *api) handleDeleteUserlist(ctx *fiber.Ctx) {
	var query map[string]int
	err := ctx.BodyParser(&query)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		logrus.Errorf("[Settings/Lists/Delete] Parsing Body: %v", err)
		return
	}

	listID, found := query["listID"]
	if !found {
		ctx.SendStatus(http.StatusBadRequest)
		return
	}

	user := ctx.Locals("user").(general.User)
	err = a.DBSession.RemoveUserList(user.ID, listID)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		logrus.Errorf("[Settings/Lists/Delete] Removing User-List: %v", err)
		return
	}

	ctx.SendStatus(http.StatusOK)
}
