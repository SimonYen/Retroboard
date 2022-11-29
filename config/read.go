package config

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
)

// 定义配置实例
var ServerConfigInstance serverConfig
var RedisConfigInstance redisConfig

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
	RedisConfigInstance.addr = viper.GetString("redis.addr")
	RedisConfigInstance.db = uint8(viper.GetUint("redis.db"))
	RedisConfigInstance.port = viper.GetUint16("redis.port")
	RedisConfigInstance.psw = viper.GetString("redis.psw")
}

// 返回地址，用于iris的Run方法
func (receiver *serverConfig) ToAddr() iris.Runner {
	return iris.Addr(fmt.Sprintf("%s:%d", receiver.addr, receiver.port))
}

// 返回redis url，用于redis的连接
func (receiver *redisConfig) ToString() string {
	return fmt.Sprintf("redis://:%s@%s:%d/%d", receiver.psw, receiver.addr, receiver.port, receiver.db)
}
