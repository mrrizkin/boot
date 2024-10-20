package logger

import (
	"github.com/mrrizkin/boot/system/config"
)

type Logger interface {
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	GetLogger() interface{}
}

func Zerolog(config *config.Config) (Logger, error) {
	return newZerolog(config)
}
