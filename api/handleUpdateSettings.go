package api

import (
	"net/http"
	"random_wikipedia/general"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (a *api) handleUpdateSettings(ctx *fiber.Ctx) {
	var nSettings general.Settings
	err := ctx.BodyParser(&nSettings)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		logrus.Errorf("[Settings/Update] Getting Body: %v", err)
		return
	}

	user := ctx.Locals("user").(general.User)

	err = a.DBSession.InsertSettings(user.ID, nSettings, true)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		logrus.Errorf("[Settings/Update] Updating Settings: %v", err)
		return
	}

	ctx.SendStatus(http.StatusOK)
}
