package api

import (
	"random_wikipedia/general"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (a *api) handleLoadUserMyself(ctx *fiber.Ctx) {
	user := ctx.Locals("user").(general.User)

	if err := ctx.JSON(user); err != nil {
		logrus.Errorf("[User/Load/Myself] Sending JSON: %v", err)
	}
}
