package welcome

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/internal/system/logger"
	"github.com/mrrizkin/boot/internal/utils"
	"github.com/mrrizkin/boot/resources/views"
)

type WelcomeController struct {
	log *logger.Logger
}

type WelcomeControllerDeps struct {
	fx.In

	Logger *logger.Logger
}

func NewWelcomeController(p WelcomeControllerDeps) (*WelcomeController, error) {
	return &WelcomeController{
		log: p.Logger,
	}, nil
}

func (a *WelcomeController) Welcome(c *fiber.Ctx) error {
	return utils.Render(c, views.Welcome("Rizki"))
}
