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
