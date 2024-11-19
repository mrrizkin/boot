package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrrizkin/boot/app/controllers"
	"github.com/mrrizkin/boot/app/providers/app"
)

func ApiRoutes(
	app *app.App,

	userController *controllers.UserController,
) {
	api := app.ApiRoutes()
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Set up v1 routes
	v1 := api.Group("/v1")

	// User routes
	v1.Get("/user", userController.UserFindAll)
	v1.Get("/user/:id", userController.UserFindByID)
	v1.Post("/user", userController.UserCreate)
	v1.Put("/user/:id", userController.UserUpdate)
	v1.Delete("/user/:id", userController.UserDelete)
}
