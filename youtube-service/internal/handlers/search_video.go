package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/youtube-service/internal/models-services"
)

// search_video handler returns all the videos matching the search query in the database in a paginated manner
func Do(c *fiber.Ctx) error {
	searchQuery := c.Query("query", "")
	if searchQuery == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "query param is required",
		})
	}

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "page query param must be an integer",
		})
	}

	return c.JSON(fiber.Map{
		"videos": models-services.(get_video-search_video).SearchVideos(searchQuery, int64(page)),
	})
}
