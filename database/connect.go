package database

import (
	"Retroboard/config"

	"github.com/gomodule/redigo/redis"
)

var RedisConn redis.Conn

// 连接到redis数据库
func Connect() {
	RC, err := redis.DialURL(config.RedisConfigInstance.ToString())
	if err != nil {
		panic(err)
	}
	RedisConn = RC
}

// 初始化数据库
// func Migrate() {
// 	//检测apps键是否存在
// }
