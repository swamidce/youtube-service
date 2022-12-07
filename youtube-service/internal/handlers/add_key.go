package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/youtube-service/internal/models-services"
)

// add_key handler adds a new API key to the database if it is valid
func Do(c *fiber.Ctx) error {
	apiKey := c.Query("key", "")
	if apiKey == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "api_key query param is required",
		})
	}

	if !add_key.IsKeyValid(apiKey) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid api key",
		})
	}

	err := add_key.InsertKey(apiKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to insert api key into the database",
		})
	}
	return c.JSON(fiber.Map{
		"message": "api key added successfully",
	})
}
