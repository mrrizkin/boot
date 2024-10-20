package server

import (
	"errors"
	"fmt"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/rs/zerolog"

	_ "github.com/joho/godotenv/autoload"

	"github.com/mrrizkin/boot/system/config"
	"github.com/mrrizkin/boot/system/validator"
	"github.com/mrrizkin/boot/third-party/logger"
)

type Server struct {
	*fiber.App

	config *config.Config
	log    logger.Logger
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

			return c.Status(code).JSON(validator.GlobalErrorResponse{
				Status: "error",
				Detail: err.Error(),
			})
		},
	})

	app.Static("/", "public")
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: logger.GetLogger().(*zerolog.Logger),
	}))
	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(idempotency.New())

	return &Server{
		App:    app,
		config: config,
		log:    logger,
	}
}

func (s *Server) Serve() error {
	s.log.Info("Server running", "port", s.config.PORT)
	return s.Listen(fmt.Sprintf(":%d", s.config.PORT))
}

func (s *Server) Stop() error {
	return s.Shutdown()
}
