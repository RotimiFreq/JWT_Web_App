package main

import (
	"hotelbooking/configs"

	"hotelbooking/routes"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.HB_routes(app)

	app.Listen(":8080")
}
