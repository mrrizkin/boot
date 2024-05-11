package logger

import (
	"io"
	"os"
	"path"
	"strings"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/mrrizkin/boot/internal/system/config"
)

type Logger struct {
	*zerolog.Logger
}

type FxLogger struct {
	Logger *Logger
}

type LoggerParams struct {
	fx.In

	Config *config.Config
}

func New(p LoggerParams) (*Logger, error) {
	var writers []io.Writer

	if p.Config.LOG_CONSOLE {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if p.Config.LOG_FILE {
		rf, err := rollingFile(p.Config)
		if err != nil {
			return nil, err
		}
		writers = append(writers, rf)
	}
	mw := io.MultiWriter(writers...)

	switch p.Config.LOG_LEVEL {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "disable":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}

	logger := zerolog.New(mw).With().Timestamp().Logger()

	logger.Info().
		Bool("fileLogging", p.Config.LOG_FILE).
		Bool("jsonLogOutput", p.Config.LOG_JSON).
		Str("logDirectory", p.Config.LOG_DIR).
		Str("fileName", p.Config.APP_NAME+".log").
		Int("maxSizeMB", p.Config.LOG_MAX_SIZE).
		Int("maxBackups", p.Config.LOG_MAX_BACKUP).
		Int("maxAgeInDays", p.Config.LOG_MAX_AGE).
		Msg("logging configured")

	return &Logger{
		Logger: &logger,
	}, nil
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

func NewFxLogger(logger *Logger) fxevent.Logger {
	return &FxLogger{
		Logger: logger,
	}
}

func (l *FxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Logger.Info().
			Str("callee", e.FunctionName).
			Str("caller", e.CallerName).
			Msg("OnStart hook executing")

	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Logger.Err(e.Err).
				Str("callee", e.FunctionName).
				Str("caller", e.CallerName).
				Msg("OnStart hook failed")
		} else {
			l.Logger.Info().
				Str("callee", e.FunctionName).
				Str("caller", e.CallerName).
				Str("runtime", e.Runtime.String()).
				Msg("OnStart hook executed")
		}
	case *fxevent.OnStopExecuting:
		l.Logger.Info().
			Str("callee", e.FunctionName).
			Str("caller", e.CallerName).
			Msg("OnStop hook executing")
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Logger.Err(e.Err).
				Str("callee", e.FunctionName).
				Str("caller", e.CallerName).
				Msg("OnStop hook failed")
		} else {
			l.Logger.Info().
				Str("callee", e.FunctionName).
				Str("caller", e.CallerName).
				Str("runtime", e.Runtime.String()).
				Msg("OnStop hook executed")
		}
	case *fxevent.Supplied:
		l.Logger.Err(e.Err).
			Str("type", e.TypeName).
			Str("module", e.ModuleName).
			Msg("supplied")
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Info().
				Str("constructor", e.ConstructorName).
				Str("module", e.ModuleName).
				Str("type", rtype).
				Msg("provided")
		}
		if e.Err != nil {
			l.Logger.Err(e.Err).
				Str("module", e.ModuleName).
				Msg("error encountered while applying options")
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Info().
				Str("decorator", e.DecoratorName).
				Str("module", e.ModuleName).
				Str("type", rtype).
				Msg("decorated")
		}
		if e.Err != nil {
			l.Logger.Err(e.Err).
				Str("module", e.ModuleName).
				Msg("error encountered while applying options")
		}
	case *fxevent.Invoking:
		// Do not log stack as it will make logs hard to read.
		l.Logger.Info().
			Str("function", e.FunctionName).
			Str("module", e.ModuleName).
			Msg("invoking")
	case *fxevent.Invoked:
		if e.Err != nil {
			l.Logger.Err(e.Err).
				Str("stack", e.Trace).
				Str("function", e.FunctionName).
				Msg("invoke failed")
		}
	case *fxevent.Stopping:
		l.Logger.Info().
			Str("signal", strings.ToUpper(e.Signal.String())).
			Msg("received signal")
	case *fxevent.Stopped:
		if e.Err != nil {
			l.Logger.Err(e.Err).
				Msg("stop failed")
		}
	case *fxevent.RollingBack:
		l.Logger.Err(e.StartErr).
			Msg("start failed, rolling back")
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.Logger.Err(e.Err).
				Msg("rollback failed")
		}
	case *fxevent.Started:
		if e.Err != nil {
			l.Logger.Err(e.Err).
				Msg("start failed")
		} else {
			l.Logger.Info().
				Msg("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.Logger.Err(e.Err).
				Msg("custom logger initialization failed")
		} else {
			l.Logger.Info().
				Str("function", e.ConstructorName).
				Msg("initialized custom fxevent.Logger")
		}
	}
}
