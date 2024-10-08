package handlers

import (
	"github.com/mrrizkin/boot/app/domains/user"
	"github.com/mrrizkin/boot/system/stypes"
)

type Handlers struct {
	*stypes.App

	userRepo    *user.Repo
	userService *user.Service
}

func New(
	app *stypes.App,
) *Handlers {
	userRepo := user.NewRepo(app.System.Database)
	userService := user.NewService(userRepo, app.Library.Hashing)

	return &Handlers{
		App: app,

		userRepo:    userRepo,
		userService: userService,
	}
}
