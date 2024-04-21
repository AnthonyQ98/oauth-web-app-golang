package main

import (
	"github.com/anthonyq98/oauth-web-app-golang/config"
	"github.com/anthonyq98/oauth-web-app-golang/controllers"
	"github.com/anthonyq98/oauth-web-app-golang/controllers/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"log"
)

func main() {
	engine := html.New("./web/template", ".html")

	app := fiber.New(fiber.Config{Views: engine})
	app.Static("/styles", "./styles")
	config.GoogleConfig()

	app.Get("/", handlers.HomeHandler)
	app.Get("/profile", handlers.ProfileHandler)
	app.Get("/login", handlers.LoginHandler)
	app.Get("/google_login", controllers.GoogleLogin)
	app.Get("/google_callback", controllers.GoogleCallback)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("Error:", err)
	}

}
