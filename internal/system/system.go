package system

import (
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/internal/domain/user"
	"github.com/mrrizkin/boot/internal/domain/welcome"
	"github.com/mrrizkin/boot/internal/model"
	"github.com/mrrizkin/boot/internal/system/config"
	"github.com/mrrizkin/boot/internal/system/database"
	"github.com/mrrizkin/boot/internal/system/logger"
	"github.com/mrrizkin/boot/internal/system/server"
	"github.com/mrrizkin/boot/internal/system/session"
)

func Run() {
	fx.New(
		fx.WithLogger(logger.NewFxLogger),
		fx.Provide(
			config.New,
			logger.New,

			model.New,
			database.New,
			session.New,

			user.NewUserRepo,
			user.NewUserAPI,
			user.NewUserService,

			welcome.NewWelcomeAPI,

			server.New,
			server.NewRoutes,
		),
		fx.Invoke(server.Serve),
	).Run()
}
