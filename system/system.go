package system

import (
	"fmt"

	"github.com/mrrizkin/boot/routes"

	"github.com/mrrizkin/boot/app/models"
	"github.com/mrrizkin/boot/system/config"
	"github.com/mrrizkin/boot/system/database"
	"github.com/mrrizkin/boot/system/server"
	"github.com/mrrizkin/boot/system/session"
	"github.com/mrrizkin/boot/system/stypes"
	"github.com/mrrizkin/boot/system/validator"
	"github.com/mrrizkin/boot/third-party/hashing"
	"github.com/mrrizkin/boot/third-party/logger"
)

func Run() {
	conf, err := config.New()
	if err != nil {
		panic(err)
	}
	log, err := logger.Zerolog(conf)
	if err != nil {
		panic(err)
	}
	sess, err := session.New(conf)
	if err != nil {
		panic(err)
	}
	defer sess.Stop()
	hash := hashing.Argon2(*conf)

	model := models.New(conf, hash)
	db, err := database.New(conf, model, log)
	if err != nil {
		panic(err)
	}
	defer db.Stop()
	err = db.Start()
	if err != nil {
		panic(err)
	}

	valid := validator.New()
	serv := server.New(conf, log)

	routes.Setup(&stypes.App{
		App: serv.App,
		System: &stypes.System{
			Logger:    log,
			Database:  db,
			Config:    conf,
			Session:   sess,
			Validator: valid,
		},
		Library: &stypes.Library{
			Hashing: hash,
		},
	}, sess)

	log.Info(fmt.Sprintf("Server is running on port %d", conf.PORT))

	if err := serv.Serve(); err != nil {
		panic(err)
	}
}
