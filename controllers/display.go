package controllers

import (
	"github.com/kataras/iris/v12"
)

func IndexController(ctx iris.Context) {
	test_list := []string{"测试", "哈哈", "牛逼"}
	ctx.ViewData("list", test_list)
	ctx.View("index.html")
}

func AboutController(ctx iris.Context) {
	ctx.View("about.html")
}
