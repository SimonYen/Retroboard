package models

import (
	"Retroboard/database"

	"github.com/gomodule/redigo/redis"
)

/*
用于管理定义的model，其实就是model从数据库里的CRUD操作
*/

type ApplicationManager struct {
	Apps       []*Application
	Name2Index map[string]int
}

func NewApplicationManager() *ApplicationManager {
	tmp := new(ApplicationManager)
	tmp.Name2Index = make(map[string]int)
	return tmp
}

// 从set中获取所有的application名字，返回获取到的数目
func (receiver *ApplicationManager) GetAllApplications() (int, error) {
	result, err := redis.Int(database.RedisConn.Do("scard", "apps"))
	if err != nil {
		return 0, err
	}
	//开始遍历set
	names, err := redis.Strings(database.RedisConn.Do("smembers", "apps"))
	if err != nil {
		return 0, err
	}
	for index, name := range names {
		tmp := new(Application)
		tmp.set(name)                              //填充值
		receiver.Apps = append(receiver.Apps, tmp) //添加到切片中
		receiver.Name2Index[name] = index          //记录应用所在的位置
	}
	return result, nil
}

// 根据名字获取App结构体，不存在的话返回nil
func (receiver *ApplicationManager) GetApplicationByName(name string) *Application {
	i, ok := receiver.Name2Index[name]
	if ok {
		return receiver.Apps[i]
	}
	return nil
}
