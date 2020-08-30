package api

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (a *api) handleLogin(ctx *fiber.Ctx) {
	var body map[string]string
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		logrus.Errorf("[Login] Parsing Body: %v", err)
		return
	}

	email, found := body["email"]
	if !found {
		ctx.SendStatus(http.StatusBadRequest)
		logrus.Errorf("[Login] Did not contain email Field")
		return
	}

	err = a.LoginSession.Login(email)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		logrus.Errorf("[Login] Could not send login: %v", err)
		return
	}

	ctx.SendStatus(http.StatusOK)
}
