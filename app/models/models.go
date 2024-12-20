package models

import (
	"go.uber.org/fx"

	"github.com/mrrizkin/boot/app/providers/database"
	"github.com/mrrizkin/boot/app/providers/logger"
	"github.com/mrrizkin/boot/config"
	"github.com/mrrizkin/boot/pkg/boot/constructor"
)

type Model struct {
	db     *database.Database
	config *config.Database
	log    *logger.Logger

	models []interface{}
}

func (m *Model) Construct() interface{} {
	return func(
		db *database.Database,
		config *config.Database,
		log *logger.Logger,
	) *Model {
		return &Model{
			db:     db,
			log:    log,
			config: config,
			models: []interface{}{
				&Permission{},
				&RolePermission{},
				&Role{},
				&User{},
			},
		}
	}
}

func (m *Model) Migrate() error {
	if len(m.models) == 0 {
		return nil
	}

	if !m.config.AUTO_MIGRATE {
		return nil
	}

	m.log.Info("migrating models", "count", len(m.models))
	return m.db.AutoMigrate(m.models...)
}

func (m *Model) Seed() error {
	return nil
}

func New() fx.Option {
	return constructor.Load(
		&Model{},
	)
}

func AutoMigrate(m *Model) error {
	return m.Migrate()
}
