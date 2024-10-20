package scheduler

import (
	"github.com/mrrizkin/boot/third-party/logger"
	robfigcron "github.com/robfig/cron"
)

type cron struct {
	*robfigcron.Cron
	log logger.Logger
}

func newCron(log logger.Logger) Scheduler {
	return &cron{
		Cron: robfigcron.New(),
		log:  log,
	}
}

func (c *cron) Add(spec string, cmd func()) {
	c.AddFunc(spec, cmd)
}

func (c *cron) Start() {
	c.log.Info("starting cron", "entries", len(c.Entries()))
	c.Cron.Start()
}
