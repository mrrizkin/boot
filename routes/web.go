package routes

import (
	"github.com/mrrizkin/boot/app/controllers"
	"github.com/mrrizkin/boot/app/providers/app"
)

func WebRoutes(
	app *app.App,

	welcomeController *controllers.WelcomeController,
) {
	router := app.WebRoutes()
	router.Get("/", welcomeController.Index)
}
