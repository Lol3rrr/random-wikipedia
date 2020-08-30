package api

import (
	"random_wikipedia/database"
	"random_wikipedia/login"
	"random_wikipedia/notifications"
	"random_wikipedia/wikipedia"

	"github.com/gofiber/fiber"
)

// NewAPI is used to obtain a new API session
func NewAPI(
	loginSession login.Session,
	notificationSession notifications.Session,
	wikipediaSession wikipedia.Session,
	dbSession database.Session) API {
	result := &api{
		LoginSession:        loginSession,
		NotificationSession: notificationSession,
		WikipediaSession:    wikipediaSession,
		DBSession:           dbSession,
	}

	app := fiber.New()

	app.Add(fiber.MethodPost, "/api/login/", result.handleLogin)
	app.Add(fiber.MethodGet, "/api/login/confirm", result.handleLoginConfirm)

	app.Add(fiber.MethodGet, "/api/publickey", result.handlePublicKey)
	app.Add(fiber.MethodPost, "/api/subscription/update", result.Auth, result.handleUpdateSubscripton)

	app.Add(fiber.MethodPost, "/api/settings/update", result.Auth, result.handleUpdateSettings)

	app.Add(fiber.MethodGet, "/api/lists/all", result.handleAllLists)

	result.App = app
	return result
}
