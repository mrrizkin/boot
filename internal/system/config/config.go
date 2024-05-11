package config

import (
	"fmt"
	"regexp"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mrrizkin/boot/internal/utils"
)

type Config struct {
	APP_NAME string
	ENV      string
	PORT     int
	PREFORK  bool

	LOG_LEVEL      string
	LOG_CONSOLE    bool
	LOG_FILE       bool
	LOG_DIR        string
	LOG_MAX_SIZE   int
	LOG_MAX_AGE    int
	LOG_MAX_BACKUP int
	LOG_JSON       bool

	DB_DRIVER   string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_SSLMODE  string

	SESSION_DRIVER string
}

var (
	// validAPP_NAME should only include alpabetical and underscore
	validAPP_NAME = regexp.MustCompile(`^[a-zA-Z_0-9]+$`)
)

func New() (*Config, error) {
	envAPP_NAME, _ := utils.EnvStr("APP_NAME", "gofast")

	if !validAPP_NAME.MatchString(*envAPP_NAME) {
		return nil, fmt.Errorf("APP_NAME is not valid")
	}

	envENV, _ := utils.EnvStr("ENV", "development")
	envPORT, _ := utils.EnvInt("PORT", 3000)
	envPREFORK, _ := utils.EnvBool("PREFORK", false)

	envLOG_LEVEL, _ := utils.EnvStr("LOG_LEVEL", "debug")
	envLOG_CONSOLE, _ := utils.EnvBool("LOG_CONSOLE", true)
	envLOG_FILE, _ := utils.EnvBool("LOG_FILE", true)
	envLOG_DIR, _ := utils.EnvStr("LOG_DIR", "./storage/log")
	envLOG_MAX_SIZE, _ := utils.EnvInt("LOG_MAX_SIZE", 50)
	envLOG_MAX_AGE, _ := utils.EnvInt("LOG_MAX_AGE", 7)
	envLOG_MAX_BACKUP, _ := utils.EnvInt("LOG_MAX_BACKUP", 20)
	envLOG_JSON, _ := utils.EnvBool("LOG_JSON", true)

	envDBDriver, _ := utils.EnvStr("DB_DRIVER", "sqlite")
	envDBHost, _ := utils.EnvStr("DB_HOST", "localhost")
	envDBPort, _ := utils.EnvInt("DB_PORT", 5432)
	envDBName, _ := utils.EnvStr("DB_NAME", "gofast")
	envDBUsername, _ := utils.EnvStr("DB_USERNAME", "gofast")
	envDBPassword, _ := utils.EnvStr("DB_PASSWORD", "gofast")
	envDBSSLMode, _ := utils.EnvStr("DB_SSLMODE", "disable")

	envSESSION_DRIVER, _ := utils.EnvStr("SESSION_DRIVER", "sqlite")

	return &Config{
		APP_NAME: *envAPP_NAME,
		ENV:      *envENV,
		PORT:     *envPORT,
		PREFORK:  *envPREFORK,

		LOG_LEVEL:      *envLOG_LEVEL,
		LOG_CONSOLE:    *envLOG_CONSOLE,
		LOG_FILE:       *envLOG_FILE,
		LOG_DIR:        *envLOG_DIR,
		LOG_MAX_SIZE:   *envLOG_MAX_SIZE,
		LOG_MAX_AGE:    *envLOG_MAX_AGE,
		LOG_MAX_BACKUP: *envLOG_MAX_BACKUP,
		LOG_JSON:       *envLOG_JSON,

		DB_DRIVER:   *envDBDriver,
		DB_HOST:     *envDBHost,
		DB_PORT:     *envDBPort,
		DB_NAME:     *envDBName,
		DB_USERNAME: *envDBUsername,
		DB_PASSWORD: *envDBPassword,
		DB_SSLMODE:  *envDBSSLMode,

		SESSION_DRIVER: *envSESSION_DRIVER,
	}, nil
}
