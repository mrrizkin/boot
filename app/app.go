// Package app provides the core application structure and functionality for the web service.
// It integrates various components such as configuration management, session handling,
// logging, rendering, and API documentation.
package app

import (
	"bytes"
	"fmt"
	"maps"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/mrrizkin/boot/system/server"
	"github.com/mrrizkin/boot/system/types"
)

// App encapsulates the application's core components and functionality.
type App struct {
	*fiber.App // Embedded Fiber application for HTTP handling

	System  *types.System  // System-wide configurations and components
	Library *types.Library // Application-specific library and utilities
}

// New initializes and returns a new App instance.
// It sets up the database, performs migrations, and seeds initial data.
func New(server *server.Server, sys *types.System) (*App, error) {
	if err := sys.Model.Migrate(sys.Database.DB); err != nil {
		return nil, fmt.Errorf("failed to run database migrations: %w", err)
	}

	if err := sys.Model.Seeds(sys.Database.DB); err != nil {
		return nil, fmt.Errorf("failed to seed database: %w", err)
	}

	return &App{
		App:     server.App,
		System:  sys,
		Library: &types.Library{},
	}, nil
}

// Config retrieves a configuration value by key.
// It provides a convenient way to access application settings.
func (a *App) Config(key string) interface{} {
	return a.System.Config.Get(key)
}

// Session retrieves the session associated with the given Fiber context.
func (a *App) Session(c *fiber.Ctx) (*session.Session, error) {
	return a.System.Session.Get(c)
}

// SessionGet retrieves a value from the session by key.
func (a *App) SessionGet(c *fiber.Ctx, key string) (interface{}, error) {
	session, err := a.Session(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}
	return session.Get(key), nil
}

// SessionSet sets a value in the session.
func (a *App) SessionSet(c *fiber.Ctx, key string, val interface{}) error {
	session, err := a.Session(c)
	if err != nil {
		return fmt.Errorf("failed to get session: %w", err)
	}
	session.Set(key, val)
	return session.Save()
}

// Log logs a message with the specified level and arguments.
// It delegates to the appropriate logging method based on the level.
func (a *App) Log(level, msg string, args ...interface{}) {
	switch strings.ToLower(level) {
	case "info":
		a.System.Logger.Info(msg, args...)
	case "warn":
		a.System.Logger.Warn(msg, args...)
	case "error":
		a.System.Logger.Error(msg, args...)
	case "fatal":
		a.System.Logger.Fatal(msg, args...)
	default:
		a.System.Logger.Info(msg, args...)
	}
}

// Render renders a template with the given name and data.
// It sets up a context with default values and merges any additional data provided.
func (a *App) Render(c *fiber.Ctx, name string, data ...fiber.Map) error {
	ctx := fiber.Map{
		"config": a.Config,
		"session": func(key string) (interface{}, error) {
			return a.SessionGet(c, key)
		},
	}

	if len(data) > 0 {
		maps.Copy(ctx, data[0])
	}

	var buf bytes.Buffer
	if err := a.System.View.Render(&buf, name, ctx); err != nil {
		return fmt.Errorf("failed to render template '%s': %w", name, err)
	}

	return c.Type("html").Send(buf.Bytes())
}

// Swagger returns a handler function for serving Swagger UI.
// It checks the environment and only serves the UI in non-production environments.
func (a *App) Swagger(swaggerPath string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		env, ok := a.Config("ENV").(string)
		if !ok {
			return &fiber.Error{
				Code:    fiber.StatusInternalServerError,
				Message: "ENV configuration is not set or invalid",
			}
		}

		if strings.ToLower(env) == "prod" || strings.ToLower(env) == "production" {
			return c.Status(fiber.StatusNotFound).Send(nil)
		}

		html := fmt.Sprintf(`<!doctype html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>Swagger API Reference - Scalar</title>
        <link rel="icon" type="image/svg+xml" href="https://docs.scalar.com/favicon.svg">
        <link rel="icon" type="image/png" href="https://docs.scalar.com/favicon.png">
    </head>
    <body>
        <script id="api-reference" data-url="%s"></script>
        <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
    </body>
</html>`, swaggerPath)

		return c.Type("html").Send([]byte(html))
	}
}
