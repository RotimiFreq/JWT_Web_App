package routes

import (
	"hotelbooking/controllers"

	"github.com/gofiber/fiber/v2"
)

func HB_routes(app *fiber.App) {
	app.Post("/Register", controllers.RegisterUser)
}