package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog"

	_ "github.com/joho/godotenv/autoload"

	"github.com/mrrizkin/boot/system/config"
	"github.com/mrrizkin/boot/system/eyoy"
	"github.com/mrrizkin/boot/system/validator"
	"github.com/mrrizkin/boot/third-party/logger"
)

type Server struct {
	*fiber.App

	config *config.Config
}

func New(config *config.Config, logger logger.Logger) *Server {
	app := fiber.New(fiber.Config{
		Prefork:               config.PREFORK,
		AppName:               config.APP_NAME,
		DisableStartupMessage: true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			title := ""

			if code < 600 {
				title = http.StatusText(code)
			} else {
				title = eyoy.StatusTitle(code)
			}

			return c.Status(code).JSON(validator.GlobalErrorResponse{
				Status: "error",
				Title:  title,
				Detail: err.Error(),
			})
		},
	})

	app.Static("/", "public")
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger.GetLogger().(*zerolog.Logger),
	}))
	app.Use(requestid.New())
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(idempotency.New())

	return &Server{
		App:    app,
		config: config,
	}
}

func (s *Server) Serve() error {
	return s.Listen(fmt.Sprintf(":%d", s.config.PORT))
}

func (s *Server) Stop() error {
	return s.Shutdown()
}
