package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/mrrizkin/boot/app"
	"github.com/mrrizkin/boot/app/handlers"
)

func WebRoutes(app *app.App, handler *handlers.Handlers) {
	ui := app.Group("/",
		csrf.New(csrf.Config{
			KeyLookup:         fmt.Sprintf("cookie:%s", app.Config("CSRF_KEY")),
			CookieName:        app.Config("CSRF_COOKIE_NAME").(string),
			CookieSameSite:    app.Config("CSRF_SAME_SITE").(string),
			CookieSecure:      app.Config("CSRF_SECURE").(bool),
			CookieSessionOnly: true,
			CookieHTTPOnly:    app.Config("CSRF_HTTP_ONLY").(bool),
			SingleUseToken:    true,
			Expiration:        time.Duration(app.Config("CSRF_EXPIRATION").(int64)) * time.Second,
			KeyGenerator:      utils.UUIDv4,
			ErrorHandler:      csrf.ConfigDefault.ErrorHandler,
			Extractor:         csrf.CsrfFromCookie(app.Config("CSRF_KEY").(string)),
			Session:           app.System.Session.Store,
			SessionKey:        "fiber.csrf.token",
			HandlerContextKey: "fiber.csrf.handler",
		}),
		cors.New(),
		helmet.New(),
	)

	ui.Get("/", func(c *fiber.Ctx) error {
		return app.Render(c, "pages/welcome")
	})
}
