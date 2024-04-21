package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"net/url"
	"strings"
)

type ProfileData struct {
	Name  string
	Email string
}

func HomeHandler(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{})
}

func ProfileHandler(c *fiber.Ctx) error {
	// Extract user data from query parameters
	userData := c.Query("data")
	if userData == "" {
		return c.SendString("User data not found in query parameters")
	}

	// Decode user data from URL-encoded string
	decodedUserData, err := url.QueryUnescape(userData)
	if err != nil {
		return c.SendString("Error decoding user data")
	}

	// Unmarshal user data from JSON
	var user map[string]interface{}
	err = json.NewDecoder(strings.NewReader(decodedUserData)).Decode(&user)
	if err != nil {
		return c.SendString("Error parsing user data")
	}

	// Render the profile page with user data
	return c.Render("profile", fiber.Map{
		"Username": user["name"],
		"Email":    user["email"],
	})
}

func LoginHandler(c *fiber.Ctx) error {
	data := ProfileData{}
	return c.Render("login", fiber.Map{
		"data": data,
	})
}
