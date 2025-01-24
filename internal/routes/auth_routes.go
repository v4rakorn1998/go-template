package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/v4rakorn1998/go-template/internal/auth"
)

func AuthRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User Routes
	api.Post("/auth", auth.Login)
	api.Post("/register", auth.Register)

}
