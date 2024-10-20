package view

import (
	"io"
	"net/http"

	"github.com/mrrizkin/boot/system/config"
)

type View interface {
	Render(w io.Writer, template string, data ...map[string]interface{}) error
}

func Jinja2(cfg *config.Config, fs http.FileSystem, directory, extension string) (View, error) {
	return newJinja2(fs, directory, extension)
}
