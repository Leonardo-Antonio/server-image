package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Home struct { }

func NewHome() *Home {
	return &Home{}
}

func (h *Home) Index(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).
		Render("index", map[string]interface{}{})
}