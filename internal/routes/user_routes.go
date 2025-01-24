package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/v4rakorn1998/go-template/internal/auth"
	"github.com/v4rakorn1998/go-template/internal/handlers"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api/users")

	// User Routes
	api.Get("/", auth.AuthMiddleware, handlers.GetUser)

}
