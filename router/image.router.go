package router

import (
	"github.com/Leonardo-Antonio/server-image/controller"
	"github.com/gofiber/fiber/v2"
)

func Image(app *fiber.App) {
	handler := controller.NewImage()
	group := app.Group("/image")
	group.Get("", handler.Image)
	group.Post("", handler.SaveImage)
}
