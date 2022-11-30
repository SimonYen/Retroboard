package config

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

// 定义配置实例
var ServerConfigInstance serverConfig
var SqliteConfigInstance sqliteConfig

func set() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(nil)
	}
}

// 将配置文件中的信息填入到相应的struct中
func Read() {
	set()
	ServerConfigInstance.addr = viper.GetString("server.addr")
	ServerConfigInstance.port = viper.GetUint16("server.port")
	SqliteConfigInstance.path = viper.GetString("sqlite.path")
}

// 返回地址，用于iris的Run方法
func (receiver *serverConfig) ToAddr() iris.Runner {
	return iris.Addr(fmt.Sprintf("%s:%d", receiver.addr, receiver.port))
}

// 返回路径，用于xorm的引擎创建
func (receiver *sqliteConfig) ToPath() string {
	return receiver.path
}
