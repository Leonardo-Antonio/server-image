package main

import (
	"github.com/Leonardo-Antonio/server-image/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/django"
	"log"
	"os"
)

type App struct {
	port string
	app *fiber.App
}

func NewAppServer(port string) *App {
	engine := django.New("./view", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	return &App{port, app}
}

func (server *App) Middlewares() {
	server.app.Use(logger.New())
}

func (server *App) Routers() {
	server.app.Static("/public", "./public")
	router.Home(server.app)
	router.Image(server.app)
}

func (server *App) Listening() {
	if len(server.port) == 0 {
		server.port = os.Getenv("PORT")
	}
	log.Fatal(server.app.Listen(":" + server.port))
}