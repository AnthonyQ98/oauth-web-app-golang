package main

import (
	"github.com/anthonyq98/oauth-web-app-golang/config"
	"github.com/anthonyq98/oauth-web-app-golang/controllers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	config.GoogleConfig()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/google_login", controllers.GoogleLogin)
	app.Get("/google_callback", controllers.GoogleCallback)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("Error:", err)
	}

}
