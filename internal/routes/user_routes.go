package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/v4rakorn1998/go-template/internal/auth"
	"github.com/v4rakorn1998/go-template/internal/handlers"
)

func UserRoutes(app *fiber.App) {
	api := app.Group("/api/users")

	// User Routes
	api.Post("/", auth.AuthMiddleware, handlers.GetUser)
	api.Post("/createUser", auth.AuthMiddleware, handlers.CreateUser)
	api.Put("/updateUser/:id", auth.AuthMiddleware, handlers.UpdateUser)
	api.Delete("/deleteUser/:id", auth.AuthMiddleware, handlers.DeleteUser)
	api.Put("/changePassword/:id", auth.AuthMiddleware, handlers.ChangePassword)
	api.Put("/resetPassword/:id", auth.AuthMiddleware, auth.CheckRoleAdmin, handlers.ResetPassword)
	api.Get("/profile", auth.AuthMiddleware, handlers.GetProfile)

}
