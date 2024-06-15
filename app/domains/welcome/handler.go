package welcome

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/app/helpers"
	"github.com/mrrizkin/boot/resources/views"
	"github.com/mrrizkin/boot/system/logger"
)

type WelcomeHandler struct {
	log *logger.Logger
}

type WelcomeHandlerDeps struct {
	fx.In

	Logger *logger.Logger
}

func NewWelcomeHandler(p WelcomeHandlerDeps) (*WelcomeHandler, error) {
	return &WelcomeHandler{
		log: p.Logger,
	}, nil
}

func (a *WelcomeHandler) Welcome(c *fiber.Ctx) error {
	return helpers.Render(c, views.Welcome("Rizki"))
}
