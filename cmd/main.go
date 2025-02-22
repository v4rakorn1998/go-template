package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/v4rakorn1998/go-template/config"
	"github.com/v4rakorn1998/go-template/internal/db"
	"github.com/v4rakorn1998/go-template/internal/routes"
)

func main() {
	// โหลด Config
	config.Load()

	db.ConnectDB()
	// สร้าง Fiber App
	app := fiber.New()

	// ตั้งค่า Routes
	routes.AuthRoutes(app)
	routes.UserRoutes(app)
	routes.CustomersRoutes(app)

	// เริ่มเซิร์ฟเวอร์
	app.Listen(":" + os.Getenv("PORT"))
	//add midder ware regis password
	//go run cmd/main.go   (for run app)
}
