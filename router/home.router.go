package router

import (
	"github.com/Leonardo-Antonio/server-image/controller"
	"github.com/gofiber/fiber/v2"
)

func Home(app *fiber.App) {
	handler := controller.NewHome()
	app.Get("/", handler.Index)
}