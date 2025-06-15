package v1

import (
	"n8n_project_go/app/analytics"
	"n8n_project_go/config"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []analytics.Apis
	result := config.DB.Find(&users)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch users"})
	}
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user analytics.Apis
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to create user"})
	}
	return c.Status(201).JSON(user)
}
