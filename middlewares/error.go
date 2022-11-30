package middlewares

import "github.com/kataras/iris/v12"

/*
错误处理中间件
*/

func PageNotFound(ctx iris.Context) {
	ctx.View("error/404.html")
}

func InternalServerError(ctx iris.Context) {
	ctx.View("error/500.html")
}
