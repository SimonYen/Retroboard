package models

/*
用于管理定义的model，其实就是model从数据库里的CRUD操作
*/

type ApplicationManager struct {
	Apps []Application
}

func NewApplicationManager() *ApplicationManager {
	return new(ApplicationManager)
}

func (receiver *ApplicationManager) GetAllApplications() {
	//从set中获取所有的application名

}
