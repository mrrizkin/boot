package logger

import (
	"io"
	"os"
	"path"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/mrrizkin/boot/system/config"
	log "github.com/rs/zerolog"
)

type zerolog struct {
	*log.Logger
}

func newZerolog(config *config.Config) (Logger, error) {
	var writers []io.Writer

	if config.LOG_CONSOLE {
		writers = append(writers, log.ConsoleWriter{Out: os.Stderr})
	}
	if config.LOG_FILE {
		rf, err := rollingFile(config)
		if err != nil {
			return nil, err
		}
		writers = append(writers, rf)
	}
	mw := io.MultiWriter(writers...)

	switch config.LOG_LEVEL {
	case "panic":
		log.SetGlobalLevel(log.PanicLevel)
	case "fatal":
		log.SetGlobalLevel(log.FatalLevel)
	case "error":
		log.SetGlobalLevel(log.ErrorLevel)
	case "warn":
		log.SetGlobalLevel(log.WarnLevel)
	case "info":
		log.SetGlobalLevel(log.InfoLevel)
	case "debug":
		log.SetGlobalLevel(log.DebugLevel)
	case "trace":
		log.SetGlobalLevel(log.TraceLevel)
	case "disable":
		log.SetGlobalLevel(log.Disabled)
	}

	logger := log.New(mw).With().Timestamp().Logger()

	logger.Info().
		Bool("fileLogging", config.LOG_FILE).
		Bool("jsonLogOutput", config.LOG_JSON).
		Str("logDirectory", config.LOG_DIR).
		Str("fileName", config.APP_NAME+".log").
		Int("maxSizeMB", config.LOG_MAX_SIZE).
		Int("maxBackups", config.LOG_MAX_BACKUP).
		Int("maxAgeInDays", config.LOG_MAX_AGE).
		Msg("logging configured")

	return &zerolog{
		Logger: &logger,
	}, nil
}

func (z *zerolog) argsParser(event *log.Event, args ...interface{}) *log.Event {
	if len(args)%2 != 0 {
		z.Warn("logger: args don't match key val")
		return event
	}

	for i := 0; i < len(args); i += 2 {
		key, ok := args[i].(string)
		if !ok {
			z.Warn("info: non-string key provided")
			continue
		}

		switch value := args[i+1].(type) {
		case bool:
			event.Bool(key, value)
		case []bool:
			event.Bools(key, value)
		case string:
			event.Str(key, value)
		case []string:
			event.Strs(key, value)
		case int:
			event.Int(key, value)
		case []int:
			event.Ints(key, value)
		case int64:
			event.Int64(key, value)
		case []int64:
			event.Ints64(key, value)
		case float32:
			event.Float32(key, value)
		case float64:
			event.Float64(key, value)
		case time.Time:
			event.Time(key, value)
		case time.Duration:
			event.Dur(key, value)
		case []byte:
			event.Bytes(key, value)
		case error:
			event.Err(value)
		default:
			event.Interface(key, value)
		}
	}

	return event
}

// usage
func (z *zerolog) Info(msg string, args ...interface{}) {
	go z.argsParser(z.Logger.Info(), args...).Msg(msg)
}

func (z *zerolog) Warn(msg string, args ...interface{}) {
	go z.argsParser(z.Logger.Warn(), args...).Msg(msg)
}

func (z *zerolog) Error(msg string, args ...interface{}) {
	go z.argsParser(z.Logger.Error(), args...).Msg(msg)
}

func (z *zerolog) Fatal(msg string, args ...interface{}) {
	go z.argsParser(z.Logger.Fatal(), args...).Msg(msg)
}

func (z *zerolog) GetLogger() interface{} {
	return z.Logger
}

func rollingFile(c *config.Config) (io.Writer, error) {
	err := os.MkdirAll(c.LOG_DIR, 0744)
	if err != nil {
		return nil, err
	}

	return &lumberjack.Logger{
		Filename:   path.Join(c.LOG_DIR, c.APP_NAME+".log"),
		MaxBackups: c.LOG_MAX_BACKUP, // files
		MaxSize:    c.LOG_MAX_SIZE,   // megabytes
		MaxAge:     c.LOG_MAX_AGE,    // days
	}, nil
}
