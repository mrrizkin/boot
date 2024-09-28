package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/mrrizkin/boot/app/handlers"
	"github.com/mrrizkin/boot/system/stypes"
)

func WebRoutes(app *stypes.App, handler *handlers.Handlers) {
	ui := app.Group("/",
		csrf.New(csrf.Config{
			KeyLookup:         "cookie:" + app.System.Config.CSRF_KEY,
			CookieName:        app.System.Config.CSRF_COOKIE_NAME,
			CookieSameSite:    app.System.Config.CSRF_SAME_SITE,
			CookieSecure:      app.System.Config.CSRF_SECURE,
			CookieSessionOnly: true,
			CookieHTTPOnly:    app.System.Config.CSRF_HTTP_ONLY,
			SingleUseToken:    true,
			Expiration:        time.Duration(app.System.Config.CSRF_EXPIRATION) * time.Second,
			KeyGenerator:      utils.UUIDv4,
			ErrorHandler:      csrf.ConfigDefault.ErrorHandler,
			Extractor:         csrf.CsrfFromCookie(app.System.Config.CSRF_KEY),
			Session:           app.System.Session.Store,
			SessionKey:        "fiber.csrf.token",
			HandlerContextKey: "fiber.csrf.handler",
		}),
		cors.New(),
	)
	ui.Get("/", func(c *fiber.Ctx) error {
		return handler.Render(c, "welcome")
	})
}
