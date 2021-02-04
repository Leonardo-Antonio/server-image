package controller

import (
	"fmt"
	"github.com/Leonardo-Antonio/server-image/helper"
	"github.com/aidarkhanov/nanoid/v2"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Image struct { }

func NewImage() *Image {
	return &Image{}
}

func (i *Image) Image(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).
		Render("image", map[string]interface{}{})
}

func (i *Image) SaveImage(ctx *fiber.Ctx) error {
	id, err := nanoid.New()
	if err != nil {
		res := helper.NewResponseJSON(helper.ERR, err.Error(), true, nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	header, err := ctx.FormFile("image")
	if err != nil {
		res := helper.NewResponseJSON(helper.ERR, err.Error(), true, nil)
		return ctx.Status(http.StatusBadRequest).JSON(res)
	}

	file, err := header.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		res := helper.NewResponseJSON(helper.ERR, err.Error(), true, nil)
		return ctx.Status(http.StatusInternalServerError).JSON(res)
	}

	fileName := id + strings.Join(strings.Split(header.Filename, " "), "-")
	err = ioutil.WriteFile(
			"public/image/" + fileName,
			fileBytes,
			os.ModePerm,
		)
	if err != nil {
		res := helper.NewResponseJSON(helper.ERR, err.Error(), true, nil)
		return ctx.Status(http.StatusInternalServerError).JSON(res)
	}
	res := helper.NewResponseJSON(
		helper.MSG,
		"image was saved successfully",
		false,
		ctx.Hostname() + "/public/image/" + fileName,
		)
	return ctx.Status(http.StatusOK).JSON(res)
}
