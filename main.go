package main

import (
	"github.com/anthonyq98/oauth-web-app-golang/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/google_login", controllers.GoogleLogin)
	app.Post("/google_callback", controllers.GoogleCallback)

	app.Listen(":8080")

}
