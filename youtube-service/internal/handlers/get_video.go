package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/youtube-service/internal/models-services"
)

// get_video handler returns all the videos in the database in a paginated manner
func Do(c *fiber.Ctx) error {

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "page query param must be an integer",
		})
	}

	videos := models-services.(get_video-search_video).GetVideos(int64(page))

	return c.JSON(fiber.Map{
		"videos": videos,
	})
}
