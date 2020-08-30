package api

import (
	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (a *api) handleAllLists(ctx *fiber.Ctx) {
	lists := a.WikipediaSession.GetLists()

	result := map[string]interface{}{
		"Lists": lists,
	}

	if err := ctx.JSON(result); err != nil {
		logrus.Errorf("[Lists/All] Sending JSON: %v", err)
	}
}
