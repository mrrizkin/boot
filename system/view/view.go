package view

import (
	"io"
	"net/http"
)

type View interface {
	Render(w io.Writer, template string, data ...map[string]interface{}) error
}

func Jinja2(fs http.FileSystem, directory, extension string) (View, error) {
	return newJinja2(fs, directory, extension)
}
