package routes

import (
	"github.com/mrrizkin/boot/app/handlers"
	"github.com/mrrizkin/boot/system/stypes"
)

func Setup(app *stypes.App) {
	handler := handlers.New(app)
	ApiRoutes(app, handler)
	WebRoutes(app, handler)
}
