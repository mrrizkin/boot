package database

import (
	"context"
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/mrrizkin/boot/internal/model"
	"github.com/mrrizkin/boot/internal/system/config"
	"github.com/mrrizkin/boot/internal/system/logger"
)

type Database struct {
	*gorm.DB
}

type DatabaseParams struct {
	fx.In

	Lc     fx.Lifecycle
	Config *config.Config
	Model  *model.Model
	Logger *logger.Logger
}

func New(p DatabaseParams) (*Database, error) {
	var (
		db  *gorm.DB
		err error
	)

	p.Logger.Info().Msg("Connecting to database")

	switch p.Config.DB_DRIVER {
	case "pgsql":
		db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", p.Config.DB_HOST, p.Config.DB_PORT, p.Config.DB_USERNAME, p.Config.DB_NAME, p.Config.DB_PASSWORD, p.Config.DB_SSLMODE)))
		if err != nil {
			return nil, err
		}
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(p.Config.DB_HOST))
		if err != nil {
			return nil, err
		}
	case "mysql":
		db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", p.Config.DB_USERNAME, p.Config.DB_PASSWORD, p.Config.DB_HOST, p.Config.DB_PORT, p.Config.DB_NAME)))
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", p.Config.DB_DRIVER)
	}

	p.Lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			if p.Config.ENV != "prod" && p.Config.ENV != "production" {
				p.Logger.Info().Msg("Migrating model")
				err = p.Model.Migrate(db)
				if err != nil {
					return err
				}

				p.Logger.Info().Msg("Seeding model")
				err = p.Model.Seeds(db)
				if err != nil {
					return err
				}
			}

			return nil
		},
		OnStop: func(context.Context) error {
			sqlDB, err := db.DB()
			if err != nil {
				return err
			}

			return sqlDB.Close()
		},
	})

	return &Database{
		DB: db,
	}, nil
}
