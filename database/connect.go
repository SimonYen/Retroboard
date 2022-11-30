package database

import (
	"Retroboard/config"
	"Retroboard/models"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

// 定义xorm引擎
var Eng *xorm.Engine

func Connect() {
	e, err := xorm.NewEngine("sqlite3", config.SqliteConfigInstance.ToPath())
	if err != nil {
		panic(err)
	}
	Eng = e
}

// 数据库迁移
func Migrate() {
	err := Eng.Sync2(new(models.Application))
	if err != nil {
		panic(err)
	}
}
