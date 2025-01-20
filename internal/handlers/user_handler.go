package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/v4rakorn1998/go-template/internal/services"
)

func GetUser(c *fiber.Ctx) error {
	user, err := services.GetUser()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}
