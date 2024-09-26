package system

import (
	"fmt"
	"net/http"

	"github.com/mrrizkin/boot/resources"
	"github.com/mrrizkin/boot/routes"

	"github.com/mrrizkin/boot/app/models"
	"github.com/mrrizkin/boot/system/config"
	"github.com/mrrizkin/boot/system/database"
	"github.com/mrrizkin/boot/system/server"
	"github.com/mrrizkin/boot/system/session"
	"github.com/mrrizkin/boot/system/stypes"
	"github.com/mrrizkin/boot/system/validator"
	"github.com/mrrizkin/boot/system/view"
	"github.com/mrrizkin/boot/third-party/hashing"
	"github.com/mrrizkin/boot/third-party/logger"
)

func Run() error {
	conf, err := config.New()
	if err != nil {
		return err
	}
	log, err := logger.Zerolog(conf)
	if err != nil {
		return err
	}
	sess, err := session.New(conf)
	if err != nil {
		return err
	}
	defer sess.Stop()
	hash := hashing.Argon2(*conf)

	model := models.New(conf, hash)
	db, err := database.New(conf, model, log)
	if err != nil {
		return err
	}
	defer db.Stop()

	valid := validator.New()
	serv := server.New(conf, log)

	v, err := view.Jinja2(http.FS(resources.Views), "/views", ".html")
	if err != nil {
		return err
	}

	routes.Setup(&stypes.App{
		App: serv.App,
		System: &stypes.System{
			Logger:    log,
			Database:  db,
			Config:    conf,
			Session:   sess,
			Validator: valid,
			View:      v,
		},
		Library: &stypes.Library{
			Hashing: hash,
		},
	}, sess)

	log.Info(fmt.Sprintf("Server is running on port %d", conf.PORT))

	if err := serv.Serve(); err != nil {
		return err
	}

	return nil
}

func MigrateDB() error {
	conf, err := config.New()
	if err != nil {
		return err
	}
	log, err := logger.Zerolog(conf)
	if err != nil {
		return err
	}
	hash := hashing.Argon2(*conf)
	model := models.New(conf, hash)
	db, err := database.New(conf, model, log)
	if err != nil {
		return err
	}

	if err := db.Migrate(); err != nil {
		return err
	}

	return nil
}

func SeedDB() error {
	conf, err := config.New()
	if err != nil {
		return err
	}
	log, err := logger.Zerolog(conf)
	if err != nil {
		return err
	}
	hash := hashing.Argon2(*conf)
	model := models.New(conf, hash)
	db, err := database.New(conf, model, log)
	if err != nil {
		return err
	}

	if err := db.Seeder(); err != nil {
		return err
	}

	return nil
}
