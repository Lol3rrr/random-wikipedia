package api

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (a *api) Auth(ctx *fiber.Ctx) {
	sessionID := ctx.Cookies("SessionID")
	if len(sessionID) <= 0 {
		ctx.SendStatus(http.StatusUnauthorized)
		return
	}

	user, err := a.DBSession.LoadUserSessionID(sessionID)
	if err != nil {
		logrus.Errorf("[Auth-Middleware] Loading User: %v", err)
		ctx.SendStatus(http.StatusBadRequest)
		return
	}

	ctx.Locals("user", user)
	ctx.Next()
}
