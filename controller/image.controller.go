package controller

import (
	"fmt"
	"github.com/Leonardo-Antonio/server-image/helper"
	"github.com/aidarkhanov/nanoid/v2"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Image struct { }

func NewImage() *Image {
	return &Image{}
}

func (i *Image) Show(ctx *fiber.Ctx) error {
	session, err := helper.Store.Get(ctx)
	if err != nil {
		panic(err)
	}
	login, _ := strconv.ParseBool(fmt.Sprintf("%v", session.Get("login")))
	fmt.Println(login)
	if !login {
		return ctx.Status(http.StatusForbidden).
			Redirect("login")
	}

	logger, err := ioutil.ReadFile("public/image/logger.txt")
	if err != nil {
		panic(err)
	}
	links := strings.Split(string(logger), "\n")
	return ctx.Status(http.StatusOK).
		Render("show", map[string]interface{}{
			"links": links[:len(links) - 1],
	})
}

func (i *Image) Image(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).
		Render("image", map[string]interface{}{})
}

func (i *Image) Login(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).
		Render("login", map[string]interface{}{})
}

func (i *Image) Verify(ctx *fiber.Ctx) error {
	if ctx.FormValue("password") != "cmcx100pre" {
		return ctx.Status(http.StatusForbidden).
			Redirect("/")
	}

	session, err := helper.Store.Get(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer session.Save()
	session.Set("login", true)

	return ctx.Status(http.StatusOK).
		Redirect("images")
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

	logger, err := os.OpenFile("public/image/logger.txt", os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	_, err = logger.WriteString(ctx.Protocol() + "://" + ctx.Hostname() + "/public/image/" + fileName + "\n")
	if err != nil {
		panic(err)
	}

	res := helper.NewResponseJSON(
		helper.MSG,
		"image was saved successfully",
		false,
		ctx.Protocol() + "://" + ctx.Hostname() + "/public/image/" + fileName,
		)
	return ctx.Status(http.StatusOK).JSON(res)
}
