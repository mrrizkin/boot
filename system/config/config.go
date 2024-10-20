package config

import (
	"path/filepath"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	APP_NAME string `env:"APP_NAME,required"`
	APP_KEY  string `env:"APP_KEY"`
	ENV      string `env:"ENV,required"`
	PORT     int    `env:"PORT,required"`
	PREFORK  bool   `env:"PREFORK,default=false"`

	STORAGE_PATH string `env:"PREFORK,default=storage"`

	LOG_LEVEL      string `env:"LOG_LEVEL,default=debug"`
	LOG_CONSOLE    bool   `env:"LOG_CONSOLE,default=true"`
	LOG_FILE       bool   `env:"LOG_FILE,default=true"`
	LOG_DIR        string `env:"LOG_DIR"`
	LOG_MAX_SIZE   int    `env:"LOG_MAX_SIZE,default=50"`
	LOG_MAX_AGE    int    `env:"LOG_MAX_AGE,default=7"`
	LOG_MAX_BACKUP int    `env:"LOG_MAX_BACKUP,default=20"`
	LOG_JSON       bool   `env:"LOG_JSON,default=true"`

	HASH_MEMORY      int `env:"HASH_MEMORY,default=64"`
	HASH_ITERATIONS  int `env:"HASH_ITERATIONS,default=10"`
	HASH_PARALLELISM int `env:"HASH_PARALLELISM,default=2"`
	HASH_SALT_LEN    int `env:"HASH_SALT_LEN,default=32"`
	HASH_KEY_LEN     int `env:"HASH_KEY_LEN,default=32"`

	SUPER_ADMIN_NAME     string `env:"SUPER_ADMIN_NAME,required"`
	SUPER_ADMIN_EMAIL    string `env:"SUPER_ADMIN_EMAIL,required"`
	SUPER_ADMIN_USERNAME string `env:"SUPER_ADMIN_USERNAME,required"`
	SUPER_ADMIN_PASSWORD string `env:"SUPER_ADMIN_PASSWORD,required"`

	DB_DRIVER   string `env:"DB_DRIVER,default=sqlite"`
	DB_HOST     string `env:"DB_HOST"`
	DB_PORT     int    `env:"DB_PORT,default=5432"`
	DB_NAME     string `env:"DB_NAME"`
	DB_USERNAME string `env:"DB_USERNAME,default=root"`
	DB_PASSWORD string `env:"DB_PASSWORD,default=root"`
	DB_SSLMODE  string `env:"DB_SSLMODE,default=disable"`

	SESSION_DRIVER    string `env:"SESSION_DRIVER,default=file"`
	SESSION_HTTP_ONLY bool   `env:"SESSION_HTTP_ONLY,default=true"`
	SESSION_SECURE    bool   `env:"SESSION_SECURE,default=true"`
	SESSION_SAME_SITE string `env:"SESSION_SAME_SITE,default=Lax"`

	CSRF_KEY         string `env:"CSRF_KEY,default=X-CSRF-Token"`
	CSRF_COOKIE_NAME string `env:"CSRF_COOKIE_NAME,default=fiber_csrf_token"`
	CSRF_SAME_SITE   string `env:"CSRF_SAME_SITE,default=Lax"`
	CSRF_SECURE      bool   `env:"CSRF_SECURE,default=false"`
	CSRF_SESSION     bool   `env:"CSRF_SESSION,default=true"`
	CSRF_HTTP_ONLY   bool   `env:"CSRF_HTTP_ONLY,default=true"`
	CSRF_EXPIRATION  int    `env:"CSRF_EXPIRATION,default=3600"`

	SWAGGER_PATH string `env:"SWAGGER_PATH,default=/docs/swagger.json"`

	cfg map[string]interface{}
}

func New() (*Config, error) {
	config := Config{
		cfg: make(map[string]interface{}),
	}
	err := load(&config)

	if config.APP_KEY == "" {
		panic("please generate the APP_KEY using pnpm generate:key")
	}

	if config.DB_HOST == "" {
		config.DB_HOST = filepath.Join(config.STORAGE_PATH, "database.db")
	}

	if config.DB_NAME == "" {
		config.DB_NAME = config.APP_NAME
	}

	if config.LOG_DIR == "" {
		config.LOG_DIR = filepath.Join(config.STORAGE_PATH, "log")
	}

	return &config, err
}

func (c *Config) Get(key string) interface{} {
	val, ok := c.cfg[key]
	if !ok {
		return nil
	}
	return val
}

func (c *Config) Set(key string, val interface{}) {
	c.cfg[key] = val
}
