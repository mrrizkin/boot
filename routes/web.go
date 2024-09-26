package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/mrrizkin/boot/app/handlers"
	"github.com/mrrizkin/boot/system/stypes"
)

func WebRoutes(app *stypes.App, handler *handlers.Handlers) {
	ui := app.Group("/", cors.New())
	ui.Get("/", func(c *fiber.Ctx) error {
		return handler.Render(c, "welcome")
	})
}
