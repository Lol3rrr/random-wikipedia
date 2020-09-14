package api

import (
	"random_wikipedia/database"
	"random_wikipedia/login"
	"random_wikipedia/notifications"
	"random_wikipedia/wikipedia"

	"github.com/gofiber/fiber/v2"
)

// API is a simple abstraction of the api itself
type API interface {
	// Start is a blocking call which will actually start the API endpoint
	Start(port int) error
}

type api struct {
	App                 *fiber.App
	LoginSession        login.Session
	NotificationSession notifications.Session
	WikipediaSession    wikipedia.Session
	DBSession           database.Session
}
