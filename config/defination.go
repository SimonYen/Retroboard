package config

// 服务器相关配置的struct
type serverConfig struct {
	addr string //监听地址
	port uint16 //监听端口号
}

// Redis数据库相关配置的struct
type redisConfig struct {
	addr string //监听地址
	port uint16 //监听端口号
	psw  string //连接密码
	db   uint8  //连接的数据库
}
