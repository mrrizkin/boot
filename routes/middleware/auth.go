package middleware

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/boot/app/handlers"
	"github.com/mrrizkin/boot/system/stypes"
)

func mockGetUserByToken(token string) (uint, string, error) {
	return 1, "session-id", nil
}

func AuthProtected(app *stypes.App, handler *handlers.Handlers) fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := app.System.Session.Get(c)
		if err != nil {
			return &fiber.Error{
				Code:    fiber.StatusInternalServerError,
				Message: "Failed to get session",
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
