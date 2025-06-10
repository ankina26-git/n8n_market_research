package v1

import (
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New()

	app.Get("/", Hello)

	return app
}
