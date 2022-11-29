package database

import (
	"Retroboard/config"

	"github.com/go-redis/redis/v9"
)

var RedisClient *redis.Client

// 连接到redis数据库
func Connect() {
	RedisClient = redis.NewClient(config.RedisConfigInstance.ToOptions())
}

// 初始化数据库
func Migrate() {

}
