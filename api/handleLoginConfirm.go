package api

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (a *api) handleLoginConfirm(ctx *fiber.Ctx) {
	email := ctx.Query("email")
	if len(email) <= 0 {
		ctx.SendStatus(http.StatusBadRequest)
		logrus.Errorf("[Login/Confirm] Missing email")
		return
	}

	password := ctx.Query("password")
	if len(password) <= 0 {
		ctx.SendStatus(http.StatusBadRequest)
		logrus.Errorf("[Login/Confirm] Missing password")
		return
	}

	sessionID, err := a.LoginSession.Authenticate(email, password)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		logrus.Errorf("[Login/Confirm] Could not authenticate: %v", err)
		return
	}

	ctx.Cookie(&fiber.Cookie{
		Name:    "SessionID",
		Value:   sessionID,
		Path:    "/",
		Expires: time.Now().Add(14 * 24 * time.Hour),
	})
	ctx.Redirect("/")
}
