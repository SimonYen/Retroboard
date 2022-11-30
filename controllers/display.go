package controllers

/*
包含了基本的页面显示控制器
*/
import (
	"Retroboard/database"
	"Retroboard/models"
	"log"

	"github.com/kataras/iris/v12"
)

func IndexController(ctx iris.Context) {
	list := make([]*models.Application, 0)
	err := database.Eng.Find(&list)
	if err != nil {
		log.Println(err)
	}
	ctx.ViewData("list", list)
	ctx.View("index.html")
}

func AddFormController(ctx iris.Context) {
	ctx.View("add.html")
}

func AboutController(ctx iris.Context) {
	ctx.View("about.html")
}
