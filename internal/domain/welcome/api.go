package welcome

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"

	"github.com/mrrizkin/gobest/internal/system/logger"
	"github.com/mrrizkin/gobest/internal/utils"
	"github.com/mrrizkin/gobest/resources/views"
)

type WelcomeAPI struct {
	log *logger.Logger
}

type WelcomeAPIParams struct {
	fx.In

	Logger *logger.Logger
}

func NewWelcomeAPI(p WelcomeAPIParams) (*WelcomeAPI, error) {
	return &WelcomeAPI{
		log: p.Logger,
	}, nil
}

func (a *WelcomeAPI) Welcome(c *fiber.Ctx) error {
	return utils.Render(c, views.Welcome("Rizki"))
}
