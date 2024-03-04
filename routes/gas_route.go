package routes

import (
	controllers "basic_server/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	//All routes related to users comes here

	app.Get("/timestamps", controllers.GetData)
}
