package stypes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mrrizkin/boot/system/config"
	"github.com/mrrizkin/boot/system/database"
	"github.com/mrrizkin/boot/system/session"
	"github.com/mrrizkin/boot/system/validator"
	"github.com/mrrizkin/boot/system/view"
	"github.com/mrrizkin/boot/third-party/hashing"
	"github.com/mrrizkin/boot/third-party/logger"
)

type Response struct {
	Title   string          `json:"title"`
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Debug   string          `json:"debug,omitempty"`
	Data    interface{}     `json:"data"`
	Meta    *PaginationMeta `json:"meta,omitempty"`
}

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type Filter struct {
	Where     string
	WhereArgs []interface{}
}

type PaginationMeta struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	PageCount int `json:"page_count"`
}

type App struct {
	*fiber.App

	System  *System
	Library *Library
}

type System struct {
	Logger    logger.Logger
	Database  *database.Database
	Config    *config.Config
	Session   *session.Session
	Validator *validator.Validator
	View      view.View
}

type Library struct {
	Hashing hashing.Hashing
}
