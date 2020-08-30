package api

import (
	"net/http"
	"random_wikipedia/general"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (a *api) handleUpdateSubscripton(ctx *fiber.Ctx) {
	var body map[string]string
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		logrus.Errorf("[Subscription/Update] Parsing body: %v", err)
		return
	}

	subscription, found := body["Subscription"]
	if !found {
		ctx.SendStatus(http.StatusBadRequest)
		return
	}

	user := ctx.Locals("user").(general.User)

	err = a.DBSession.InsertSubscription(user.ID, subscription, true)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		logrus.Errorf("[Subscription/Update] Updating subscription: %v", err)
		return
	}

	ctx.SendStatus(http.StatusOK)
}
