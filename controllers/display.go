package controllers

import (
	"Retroboard/database"
	"Retroboard/models"

	"github.com/kataras/iris/v12"
)

func IndexController(ctx iris.Context) {
	list := make([]*models.Application, 0)
	err := database.Eng.Find(&list)
	if err != nil {

	}
	ctx.View("index.html")
}

func AboutController(ctx iris.Context) {
	ctx.View("about.html")
}
