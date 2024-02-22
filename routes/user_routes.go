package routes

import (
	"main/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Main"})
	})

	app.Post("/user", controllers.CreateUser)

	app.Get("/user/:userId", controllers.GetAUser)

	app.Put("/user/:userId", controllers.EditAUser)

	app.Delete("/user/:userId", controllers.DeleteAUser)

	app.Get("/users", controllers.GetAllUsers)

}

