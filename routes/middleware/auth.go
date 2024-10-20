package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/boot/app"
	"github.com/mrrizkin/boot/app/handlers"
)

func AuthProtected(app *app.App, handler *handlers.Handlers) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := app.Session(c)
		if err != nil {
			return &fiber.Error{
				Code:    fiber.StatusInternalServerError,
				Message: fmt.Sprintf("failed to get session: %s", err),
			}
		}

		uid, ok := session.Get("uid").(uint)
		if !ok {
			return &fiber.Error{
				Code:    fiber.StatusUnauthorized,
				Message: "Unauthorized",
			}
		}

		sid, ok := session.Get("sid").(string)
		if !ok {
			return &fiber.Error{
				Code:    fiber.StatusUnauthorized,
				Message: "Unauthorized",
			}
		}

		c.Locals("uid", uid)
		c.Locals("sid", sid)

		return c.Next()
	}
}
