package bootstrap

import (
	"fmt"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"github.com/mrrizkin/boot/app/console"
	"github.com/mrrizkin/boot/app/controllers"
	"github.com/mrrizkin/boot/app/middleware"
	"github.com/mrrizkin/boot/app/models"
	"github.com/mrrizkin/boot/app/providers"
	"github.com/mrrizkin/boot/app/providers/app"
	"github.com/mrrizkin/boot/app/providers/logger"
	"github.com/mrrizkin/boot/app/providers/scheduler"
	"github.com/mrrizkin/boot/app/repositories"
	"github.com/mrrizkin/boot/app/services"
	"github.com/mrrizkin/boot/config"
	"github.com/mrrizkin/boot/routes"
)

func App() *fx.App {
	return fx.New(
		config.New(),
		controllers.New(),
		middleware.New(),
		models.New(),
		providers.New(),
		repositories.New(),
		services.New(),

		fx.Invoke(
			app.Boot,
			console.Schedule,
			models.AutoMigrate,
			routes.ApiRoutes,
			routes.WebRoutes,
			serveHTTP,
			startScheduler,
		),

		fx.WithLogger(useLogger),
	)
}

func serveHTTP(app *app.App, cfg *config.App, log *logger.Logger) error {
	log.Info("starting server", "port", cfg.PORT)
	return app.Listen(fmt.Sprintf(":%d", cfg.PORT))
}

func startScheduler(scheduler *scheduler.Scheduler) {
	scheduler.Start()
}

func useLogger(logger *logger.Logger) fxevent.Logger {
	return logger
}
