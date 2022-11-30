package main

import (
	"Retroboard/config"
	"Retroboard/controllers"
	"Retroboard/views"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	//读取配置文件
	config.Read()
	//连接数据库

	app := iris.New()
	// recover 中间件从任何异常中恢复，如果有异常，则写入500状态码（服务器内部错误）。
	app.Use(recover.New())

	//处理静态文件
	app.HandleDir("/assets", iris.Dir("./views/assets"))

	requestLogger := logger.New(logger.Config{
		// 是否开启状态码
		Status: true,
		// 是否记录远程IP地址
		IP: true,
		// 是否呈现HTTP谓词
		Method: true,
		// 是否记录请求路径
		Path: true,
		// 是否开启查询追加
		Query: true,
	})
	app.Use(requestLogger)
	//模板引擎初始化
	views.Init(true)
	//注册模板引擎
	app.RegisterView(views.Engine)

	//路由
	app.Get("/", controllers.IndexController)
	app.Get("/about", controllers.AboutController)

	//运行
	app.Run(config.ServerConfigInstance.ToAddr())
}
