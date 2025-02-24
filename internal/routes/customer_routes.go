package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/v4rakorn1998/go-template/internal/auth"
	"github.com/v4rakorn1998/go-template/internal/handlers"
)

func CustomersRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User Routes
	api.Post("/customers", auth.AuthMiddleware, handlers.GetCustomersList)
	api.Post("/createCustomer", auth.AuthMiddleware, handlers.CreateCustomer)
	api.Put("/updateCustomer/:id", auth.AuthMiddleware, handlers.UpdateCustomer)
	api.Delete("/deleteCustomer/:id", auth.AuthMiddleware, handlers.DeleteCustomer)

}
