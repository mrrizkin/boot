package routes

import (
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/mrrizkin/boot/app/handlers"
	"github.com/mrrizkin/boot/routes/middleware"
	"github.com/mrrizkin/boot/system/stypes"
)

func ApiRoutes(app *stypes.App, handler *handlers.Handlers) {
	api := app.Group("/api", cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, finteligo-api-token",
	}), middleware.AuthProtected(handler.App, handler))

	v1 := api.Group("/v1")
	v1.Get("/user", handler.UserFindAll)
	v1.Get("/user/:id", handler.UserFindByID)
	v1.Post("/user", handler.UserCreate)
	v1.Put("/user/:id", handler.UserUpdate)
	v1.Delete("/user/:id", handler.UserDelete)
}
