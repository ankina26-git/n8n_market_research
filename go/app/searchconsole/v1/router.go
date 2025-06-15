package v1

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router) {
	router.Get("/", GetUsers)
	router.Post("/", CreateUser)
}
