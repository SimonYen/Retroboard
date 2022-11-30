package config

// 服务器相关配置的struct
type serverConfig struct {
	addr string //监听地址
	port uint16 //监听端口号
}

// 数据库相关配置的struct
type sqliteConfig struct {
	path string //db文件存放的目录
}
