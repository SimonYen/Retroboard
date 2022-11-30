package controllers

import (
	"Retroboard/database"
	"Retroboard/models"
	"log"

	"github.com/kataras/iris/v12"
)

/*
包含了application的创建，修改，删除控制器
*/

func CreateApplicationController(ctx iris.Context) {
	app := new(models.Application)
	//获取表格数据
	app.Name = ctx.PostValue("app_name")
	app.Description = ctx.PostValue("app_desc")
	app.URL = ctx.PostValue("app_url")
	//数据库创建新的row
	_, err := database.Eng.Insert(app)
	if err != nil {
		log.Println(err)
		ctx.Redirect("/", iris.StatusInternalServerError)
	}
	ctx.Redirect("/")
}
