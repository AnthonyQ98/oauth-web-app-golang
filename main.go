package main

import (
	"github.com/Siddheshk02/go-oauth2/controllers" //imoprting the controllers package
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/google_login", controllers.GoogleLogin)
	app.Post("/google_callback", controllers.GoogleCallback)

	app.Listen(":8080")

}
