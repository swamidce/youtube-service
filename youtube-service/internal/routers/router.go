package router

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/youtube-service/internal/handlers"
)


// SetRoutes initializes middlewares and creates
// routes for all the endpoints
func SetRoutes(app *fiber.App) {
	app.Get("/get_video", func(c *fiber.Ctx) error {
		return get_video.Do(c)
	})

	app.Get("/search_video", func(c *fiber.Ctx) error {
		return search_video.Do(c)
	})

	app.Post("/add_key", func(c *fiber.Ctx) error {
		return add_key.Do(c)
	})
}
