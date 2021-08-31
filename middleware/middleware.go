package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

/// Limit post requests to the official client
func CustomeMiddleware() func(*fiber.Ctx) error {

	return func(c *fiber.Ctx) error {
		if c.Method() == "POST" || c.Method() == "DELETE" {
			apiKey, envExists := os.LookupEnv("API_KEY")
			if !envExists {
				return c.Status(500).SendString("Could not fetch env vars, call 100lvlmaster")
			}
			key := c.Get("Authorization")
			if apiKey != key {
				return c.Status(400).SendString("Your keys are incorrect peasant ✨✨")
			}
		}
		return c.Next()
	}
}
