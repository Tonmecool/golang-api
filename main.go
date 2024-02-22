package main

import (
	"main/configs"
	"main/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

	//Подключение к бд
	configs.ConnectDB()

	//routes
	routes.UserRoute(app)

    app.Listen(":3000")
}

