package system

import (
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/app/domains/user"
	"github.com/mrrizkin/boot/app/domains/welcome"
	"github.com/mrrizkin/boot/app/models"

	"github.com/mrrizkin/boot/system/config"
	"github.com/mrrizkin/boot/system/database"
	"github.com/mrrizkin/boot/system/logger"
	"github.com/mrrizkin/boot/system/server"
	"github.com/mrrizkin/boot/system/session"
)

func Run() {
	fx.New(
		fx.WithLogger(logger.NewFxLogger),
		fx.Provide(
			config.New,
			logger.New,

			models.New,
			database.New,
			session.New,

			user.NewUserRepo,
			user.NewUserHandler,
			user.NewUserService,

			welcome.NewWelcomeHandler,

			server.New,
			server.NewRoutes,
		),
		fx.Invoke(server.Serve),
	).Run()
}
