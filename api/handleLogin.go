package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (a *api) handleLogin(ctx *fiber.Ctx) error {
	var body map[string]string
	err := ctx.BodyParser(&body)
	if err != nil {
		logrus.Errorf("[Login] Parsing Body: %v", err)
		return ctx.SendStatus(http.StatusBadRequest)
	}

	email, found := body["email"]
	if !found {
		logrus.Errorf("[Login] Did not contain email Field")
		return ctx.SendStatus(http.StatusBadRequest)
	}

	err = a.LoginSession.Login(email)
	if err != nil {
		logrus.Errorf("[Login] Could not send login: %v", err)
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.SendStatus(http.StatusOK)
}
