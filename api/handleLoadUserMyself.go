package api

import (
	"random_wikipedia/general"

	"github.com/gofiber/fiber/v2"
)

func (a *api) handleLoadUserMyself(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(general.User)

	return ctx.JSON(user)
}
