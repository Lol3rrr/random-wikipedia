package api

import (
	"github.com/gofiber/fiber/v2"
)

func (a *api) handlePublicKey(ctx *fiber.Ctx) error {
	result := map[string]string{
		"PublicKey": a.NotificationSession.GetPublicKey(),
	}

	return ctx.JSON(result)
}
