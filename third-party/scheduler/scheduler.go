package scheduler

import "github.com/mrrizkin/boot/third-party/logger"

type Scheduler interface {
	Add(spec string, cmd func())
	Start()
}

func Cron(log logger.Logger) Scheduler {
	return newCron(log)
}
