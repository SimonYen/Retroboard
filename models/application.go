package models

// web应用结构体，包含了名字，url，描述
type Application struct {
	Id          int64
	Name        string `xorm:"varchar(50)"`
	URL         string `xorm:"varchar(200)"`
	Description string `xorm:"varchar(250)"`
}
