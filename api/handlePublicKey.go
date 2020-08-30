package api

import (
	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (a *api) handlePublicKey(ctx *fiber.Ctx) {
	result := map[string]string{
		"PublicKey": a.NotificationSession.GetPublicKey(),
	}

	err := ctx.JSON(result)
	if err != nil {
		logrus.Errorf("Could not send JSON response: %v", err)
	}
}
