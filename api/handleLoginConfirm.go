package api

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (a *api) handleLoginConfirm(ctx *fiber.Ctx) error {
	email := ctx.Query("email")
	if len(email) <= 0 {
		logrus.Errorf("[Login/Confirm] Missing email")
		return ctx.SendStatus(http.StatusBadRequest)
	}

	password := ctx.Query("password")
	if len(password) <= 0 {
		logrus.Errorf("[Login/Confirm] Missing password")
		return ctx.SendStatus(http.StatusBadRequest)
	}

	sessionID, err := a.LoginSession.Authenticate(email, password)
	if err != nil {
		logrus.Errorf("[Login/Confirm] Could not authenticate: %v", err)
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "SessionID",
		Value:   sessionID,
		Path:    "/",
		Expires: time.Now().Add(14 * 24 * time.Hour),
	})
	return ctx.Redirect("/")
}
