package system

import (
	"net/http"

	"github.com/mrrizkin/boot/app"
	"github.com/mrrizkin/boot/resources"
	"github.com/mrrizkin/boot/routes"

	"github.com/mrrizkin/boot/app/models"
	"github.com/mrrizkin/boot/system/config"
	"github.com/mrrizkin/boot/system/database"
	"github.com/mrrizkin/boot/system/server"
	"github.com/mrrizkin/boot/system/session"
	"github.com/mrrizkin/boot/system/types"
	"github.com/mrrizkin/boot/system/validator"
	"github.com/mrrizkin/boot/system/view"
	"github.com/mrrizkin/boot/third-party/hashing"
	"github.com/mrrizkin/boot/third-party/logger"
	"github.com/mrrizkin/boot/third-party/scheduler"
)

func Run() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	log, err := logger.Zerolog(cfg)
	if err != nil {
		return err
	}

	hash := hashing.Argon2(cfg)

	valid := validator.New()
	serv := server.New(cfg, log)
	model := models.New(cfg, hash)

	db, err := database.New(cfg, model, log)
	if err != nil {
		return err
	}
	defer db.Stop()

	sess, err := session.New(cfg)
	if err != nil {
		return err
	}
	defer sess.Stop()

	v, err := view.Jinja2(cfg, http.FS(resources.Views), "/views", ".html")
	if err != nil {
		return err
	}

	sys := types.System{
		Logger:    log,
		Database:  db,
		Config:    cfg,
		Session:   sess,
		Validator: valid,
		View:      v,
		Hashing:   hash,
		Model:     model,
	}

	app, err := app.New(serv, &sys)
	if err != nil {
		return err
	}

	schedule := scheduler.Cron(log)
	routes.Setup(app)
	app.Schedule(schedule)
	schedule.Start()

	if err := serv.Serve(); err != nil {
		return err
	}

	return nil
}
