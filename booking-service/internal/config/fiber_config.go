package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewFiberConfig() *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName: "Shedule Service",
		Prefork: false,
	})

	// Menambahkan middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", 
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, PATCH",
	}))

	return app
}
