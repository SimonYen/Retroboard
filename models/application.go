package models

import (
	"Retroboard/database"

	"github.com/gomodule/redigo/redis"
)

// web应用结构体，包含了名字，url，描述
type Application struct {
	Name        string
	URL         string
	Description string
}

// 创建一个结构体放入redis中
func (receiver *Application) CreateOne(name, url, desc string) bool {
	//首先查看
}

// 用于根据应用名从redis找到相关信息并填入，这个就不返回error了，因为出错了没关系，用空值代表就行了
func (receiver *Application) set(name string) {
	receiver.Name = name
	receiver.getField(1)
	receiver.getField(2)
}

// 获取相应的field
func (receiver *Application) getField(option int) {
	field := ""
	if option == 1 {
		field = "url"
	} else {
		field = "desc"
	}
	result, _ := redis.String(database.RedisConn.Do("hget", receiver.Name+"-app", field))
	if option == 1 {
		receiver.URL = result
	} else {
		receiver.Description = result
	}
}
