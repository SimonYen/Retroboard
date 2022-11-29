package controllers

import (
	"github.com/kataras/iris/v12"
)

func IndexController(ctx iris.Context) {
	ctx.View("index.html")
}

func AboutController(ctx iris.Context) {
	ctx.View("about.html")
}
