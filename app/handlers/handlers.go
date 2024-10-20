// Package handlers provides the HTTP request handlers for the application.
// This file defines the structure and initialization of the handlers,
// which serve as an interface between incoming HTTP requests and the
// application's domain business logic.
//
// Note: The actual handler methods (e.g., UserCreate, UserFindAll, etc.)
// are defined in separate files within this package. Those files contain
// the specific logic for translating HTTP requests into calls to the
// domain services and formatting the responses.
package handlers

import (
	"github.com/mrrizkin/boot/app"
	"github.com/mrrizkin/boot/app/domains/user"
)

// Handlers encapsulates all HTTP request handlers and their dependencies.
type Handlers struct {
	*app.App

	userRepo    *user.Repo
	userService *user.Service
}

// New creates and returns a new instance of Handlers.
// It initializes the necessary repositories and services,
// setting up the dependencies for all HTTP request handlers.
func New(
	app *app.App,
) *Handlers {
	userRepo := user.NewRepo(app.System.Database)
	userService := user.NewService(userRepo, app.System.Hashing)

	return &Handlers{
		App: app,

		userRepo:    userRepo,
		userService: userService,
	}
}
