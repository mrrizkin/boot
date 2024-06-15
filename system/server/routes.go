package server

import (
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/app/domains/user"
	"github.com/mrrizkin/boot/app/domains/welcome"
)

type Routes struct {
	*Server

	UserHandler    *user.UserHandler
	WelcomeHandler *welcome.WelcomeHandler
}

type RoutesDeps struct {
	fx.In

	Server         *Server
	UserHandler    *user.UserHandler
	WelcomeHandler *welcome.WelcomeHandler
}

func NewRoutes(p RoutesDeps) *Routes {
	return &Routes{
		Server:         p.Server,
		UserHandler:    p.UserHandler,
		WelcomeHandler: p.WelcomeHandler,
	}
}

func (route *Routes) setup() {
	route.Get("/", route.WelcomeHandler.Welcome)

	api := route.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")
	user.Get("/", route.UserHandler.FindAll)
	user.Get("/:id", route.UserHandler.FindByID)
	user.Post("/", route.UserHandler.Create)
	user.Put("/:id", route.UserHandler.Update)
	user.Delete("/:id", route.UserHandler.Delete)
}
