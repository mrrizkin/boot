package main

import (
	"go.uber.org/fx"

	"github.com/mrrizkin/gobest/internal/domain/user"
	"github.com/mrrizkin/gobest/internal/domain/welcome"
	"github.com/mrrizkin/gobest/internal/model"
	"github.com/mrrizkin/gobest/internal/system/config"
	"github.com/mrrizkin/gobest/internal/system/database"
	"github.com/mrrizkin/gobest/internal/system/logger"
	"github.com/mrrizkin/gobest/internal/system/server"
	"github.com/mrrizkin/gobest/internal/system/session"
)

func main() {
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
