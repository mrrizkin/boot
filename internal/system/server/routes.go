package server

import (
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/internal/domain/user"
	"github.com/mrrizkin/boot/internal/domain/welcome"
)

type Routes struct {
	*Server

	UserController    *user.UserController
	WelcomeController *welcome.WelcomeController
}

type RoutesDeps struct {
	fx.In

	Server            *Server
	UserController    *user.UserController
	WelcomeController *welcome.WelcomeController
}

func NewRoutes(p RoutesDeps) *Routes {
	return &Routes{
		Server:            p.Server,
		UserController:    p.UserController,
		WelcomeController: p.WelcomeController,
	}
}

func (route *Routes) setup() {
	route.Get("/", route.WelcomeController.Welcome)

	api := route.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")
	user.Get("/", route.UserController.FindAll)
	user.Get("/:id", route.UserController.FindByID)
	user.Post("/", route.UserController.Create)
	user.Put("/:id", route.UserController.Update)
	user.Delete("/:id", route.UserController.Delete)
}
