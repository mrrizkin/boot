package routes

import (
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/mrrizkin/boot/app"
	"github.com/mrrizkin/boot/app/handlers"
	"github.com/mrrizkin/boot/routes/middleware"
)

// @title						Boot API
// @version					1.0
// @description				Boot API provides a comprehensive set of endpoints for managing user data and related operations.
// @termsOfService				https://www.example.com/terms/
// @contact.name				API Support Team
// @contact.url				https://www.example.com/support
// @contact.email				support@example.com
// @license.name				MIT License
// @license.url				https://opensource.org/licenses/MIT
// @host						api.example.com
// @BasePath					/api/v1
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						boot-api-token
// @externalDocs.description	OpenAPI Specification
// @externalDocs.url			https://swagger.io/specification/
func ApiRoutes(app *app.App, handler *handlers.Handlers) {
	// Set up the swagger ui
	app.Get("/api/v1/docs/swagger", app.Swagger(app.Config("SWAGGER_PATH").(string)))

	// Set up the main API group with CORS and authentication middleware
	api := app.Group("/api", cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, boot-api-token",
	}), middleware.AuthProtected(handler.App, handler))

	// Set up v1 routes
	v1 := api.Group("/v1")

	// User routes
	v1.Get("/user", handler.UserFindAll)
	v1.Get("/user/:id", handler.UserFindByID)
	v1.Post("/user", handler.UserCreate)
	v1.Put("/user/:id", handler.UserUpdate)
	v1.Delete("/user/:id", handler.UserDelete)

	// TODO: Add more route groups and endpoints as needed
}
