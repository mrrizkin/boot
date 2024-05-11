package server

import (
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/internal/domain/user"
	"github.com/mrrizkin/boot/internal/domain/welcome"
)

type Routes struct {
	*Server

	UserAPI    *user.UserAPI
	WelcomeAPI *welcome.WelcomeAPI
}

type RoutesParams struct {
	fx.In

	Server     *Server
	UserAPI    *user.UserAPI
	WelcomeAPI *welcome.WelcomeAPI
}

func NewRoutes(p RoutesParams) *Routes {
	return &Routes{
		Server:     p.Server,
		UserAPI:    p.UserAPI,
		WelcomeAPI: p.WelcomeAPI,
	}
}

func (r *Routes) setup() {
	r.Get("/", r.WelcomeAPI.Welcome)

	api := r.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")
	user.Get("/", r.UserAPI.FindAll)
	user.Get("/:id", r.UserAPI.FindByID)
	user.Post("/", r.UserAPI.Create)
	user.Put("/:id", r.UserAPI.Update)
	user.Delete("/:id", r.UserAPI.Delete)
}
