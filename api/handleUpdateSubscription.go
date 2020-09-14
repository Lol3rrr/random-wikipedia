package api

import (
	"net/http"
	"random_wikipedia/general"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (a *api) handleUpdateSubscripton(ctx *fiber.Ctx) error {
	var body map[string]string
	err := ctx.BodyParser(&body)
	if err != nil {
		logrus.Errorf("[Subscription/Update] Parsing body: %v", err)
		return ctx.SendStatus(http.StatusBadRequest)
	}

	subscription, found := body["Subscription"]
	if !found {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	user := ctx.Locals("user").(general.User)

	err = a.DBSession.InsertSubscription(user.ID, subscription, true)
	if err != nil {
		logrus.Errorf("[Subscription/Update] Updating subscription: %v", err)
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.SendStatus(http.StatusOK)
}
