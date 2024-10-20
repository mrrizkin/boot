package routes

import (
	"github.com/mrrizkin/boot/app"
	"github.com/mrrizkin/boot/app/handlers"
)

func Setup(app *app.App) {
	handler := handlers.New(app)
	ApiRoutes(app, handler)
	WebRoutes(app, handler)
}
