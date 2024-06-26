package server

import (
	"context"
	"fmt"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/system/config"
	"github.com/mrrizkin/boot/system/logger"
)

type Server struct {
	*fiber.App
}

type ServerDeps struct {
	fx.In

	Config *config.Config
	Logger *logger.Logger
}

func New(p ServerDeps) *Server {
	app := fiber.New(fiber.Config{
		Prefork:               p.Config.PREFORK,
		DisableStartupMessage: true,
		AppName:               p.Config.APP_NAME,
	})

	app.Static("/", "public")

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: p.Logger.Logger,
	}))
	app.Use(requestid.New())
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(idempotency.New())

	return &Server{
		App: app,
	}
}

type ServeDeps struct {
	fx.In

	Lc     fx.Lifecycle
	Config *config.Config
	Server *Server
	Routes *Routes
}

func Serve(p ServeDeps) {
	p.Routes.setup()

	p.Lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go p.Server.Listen(fmt.Sprintf(":%d", p.Config.PORT))
			return nil
		},
		OnStop: func(context.Context) error {
			return p.Server.Shutdown()
		},
	})
}
