package api

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func getSessionID(ctx *fiber.Ctx) (string, error) {
	sessionCookie := ctx.Cookies("SessionID")
	if len(sessionCookie) > 0 {
		return sessionCookie, nil
	}

	sessionHeader := ctx.Get("x-session")
	if len(sessionHeader) > 0 {
		return sessionHeader, nil
	}

	return "", errors.New("Could not find SessionID")
}

func (a *api) Auth(ctx *fiber.Ctx) error {
	sessionID, err := getSessionID(ctx)
	if err != nil {
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	user, err := a.DBSession.LoadUserSessionID(sessionID)
	if err != nil {
		logrus.Errorf("[Auth-Middleware] Loading User: %v", err)
		return ctx.SendStatus(http.StatusBadRequest)
	}

	ctx.Locals("user", user)
	return ctx.Next()
}
