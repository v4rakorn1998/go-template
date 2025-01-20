package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/v4rakorn1998/go-template/internal/handlers"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	// User Routes
	api.Get("/users", handlers.GetUser)
}
