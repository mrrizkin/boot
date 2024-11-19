package controllers

import (
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/pkg/boot/constructor"
)

func New() fx.Option {
	return constructor.Load(
		&WelcomeController{},
		&UserController{},
	)
}
