package api

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func getSessionID(ctx *fiber.Ctx) (string, error) {
	sessionCookie := ctx.Cookies("SessionID")
	if len(sessionCookie) > 0 {
		return sessionCookie, nil
	}

	sessionHeader := ctx.Get("x-session")
	if len(sessionCookie) > 0 {
		return sessionHeader, nil
	}

	return "", errors.New("Could not find SessionID")
}

func (a *api) Auth(ctx *fiber.Ctx) {
	sessionID, err := getSessionID(ctx)
	if err != nil {
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
