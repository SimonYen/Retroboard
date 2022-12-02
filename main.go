package main

import (
	"Retroboard/config"
	"Retroboard/controllers"
	"Retroboard/database"
	"Retroboard/middlewares"
	"Retroboard/views"
	"flag"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	//定义命令行参数
	addressPtr := flag.String("addr", "0.0.0.0:10001", "Server address, default value: 0.0.0.0:10001")
	pathPtr := flag.String("path", "./board.db", "Sqlite3 database path, default value: ./board.db")
	configPtr := flag.String("config", "", "Config file directoty, default value: \"\" , means you do not use config file.\n If you want to use config file, pls rename the config file as config.toml")

	flag.Parse()
	useConfig := false
	if *configPtr != "" {
		useConfig = true
	}

	//读取配置文件,如果设置了
	if useConfig {
		config.Read(*configPtr)
	}
	//连接数据库
	database.Connect(useConfig, *pathPtr)
	//迁移数据库
	database.Migrate()
	app := iris.New()
	// recover 中间件从任何异常中恢复，如果有异常，则写入500状态码（服务器内部错误）。
	app.Use(recover.New())
	app.OnErrorCode(iris.StatusNotFound, middlewares.PageNotFound)
	app.OnErrorCode(iris.StatusInternalServerError, middlewares.InternalServerError)
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
	app.Get("/add", controllers.AddFormController)
	app.Post("/create", controllers.CreateApplicationController)
	app.Get("/delete/{id:int64}", controllers.DeleteApplicationController)

	//运行
	if useConfig {
		app.Run(config.ServerConfigInstance.ToAddr())
	} else {
		app.Run(iris.Addr(*addressPtr))
	}
}
