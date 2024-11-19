package providers

import (
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/pkg/boot/constructor"

	"github.com/mrrizkin/boot/app/providers/app"
	"github.com/mrrizkin/boot/app/providers/asset"
	"github.com/mrrizkin/boot/app/providers/cache"
	"github.com/mrrizkin/boot/app/providers/database"
	"github.com/mrrizkin/boot/app/providers/hashing"
	"github.com/mrrizkin/boot/app/providers/logger"
	"github.com/mrrizkin/boot/app/providers/scheduler"
	"github.com/mrrizkin/boot/app/providers/session"
	"github.com/mrrizkin/boot/app/providers/validator"
	"github.com/mrrizkin/boot/app/providers/view"
)

func New() fx.Option {
	return constructor.Load(
		&app.App{},
		&asset.Asset{},
		&cache.Cache{},
		&database.Database{},
		&hashing.Hashing{},
		&logger.Logger{},
		&scheduler.Scheduler{},
		&session.Session{},
		&validator.Validator{},
		&view.View{},
	)
}
