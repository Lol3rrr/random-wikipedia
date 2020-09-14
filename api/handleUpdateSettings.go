package api

import (
	"net/http"
	"random_wikipedia/general"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (a *api) handleUpdateSettings(ctx *fiber.Ctx) error {
	var nSettings general.Settings
	err := ctx.BodyParser(&nSettings)
	if err != nil {
		logrus.Errorf("[Settings/Update] Getting Body: %v", err)
		return ctx.SendStatus(http.StatusBadRequest)
	}

	user := ctx.Locals("user").(general.User)

	err = a.DBSession.InsertSettings(user.ID, nSettings, true)
	if err != nil {
		logrus.Errorf("[Settings/Update] Updating Settings: %v", err)
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.SendStatus(http.StatusOK)
}
