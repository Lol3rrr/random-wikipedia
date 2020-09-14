package api

import (
	"github.com/gofiber/fiber/v2"
)

func (a *api) handleAllLists(ctx *fiber.Ctx) error {
	lists := a.WikipediaSession.GetLists()

	result := map[string]interface{}{
		"Lists": lists,
	}

	return ctx.JSON(result)
}
