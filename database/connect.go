package database

import (
	"Retroboard/config"
	"Retroboard/models"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

// 定义xorm引擎
var Eng *xorm.Engine

func Connect(useConfigFile bool, path string) {
	arg := ""
	if useConfigFile {
		arg = config.SqliteConfigInstance.ToPath()
	} else {
		arg = path
	}
	e, err := xorm.NewEngine("sqlite3", arg)
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
