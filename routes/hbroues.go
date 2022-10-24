package routes

import (
	"hotelbooking/controllers"
	"hotelbooking/Authmiddleware"

	"github.com/gofiber/fiber/v2"
)

func HB_routes(app *fiber.App) {
	app.Post("/Register", controllers.RegisterUser)
	app.Post("/Login", controllers.Login)

	app.Use(authmiddleware.Is_User_Authenticated)

	app.Get("/GetUser", controllers.RegisterUser)
	


}