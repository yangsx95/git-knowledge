package app

import (
	"git-knowledge/api"
	"git-knowledge/dao"
	"git-knowledge/provider"
)

// Dao 应用程序组件容器，所有Dao组件都需要注册到该文件中
// 注意，要按照顺序依次注入
type Dao struct {
	UserDao dao.UserDao
}

func initDao(b *BootStrap) *Dao {
	d := Dao{}
	d.UserDao = dao.NewUserDao(b.db)
	return &d
}

// ServiceProvider Dao的上一层，API实际提供方
type ServiceProvider struct {
	LoginService api.LoginService
}

func initServiceProvider(b *BootStrap) *ServiceProvider {
	s := ServiceProvider{}
	s.LoginService = provider.NewLoginService(b.Dao.UserDao)
	return &s
}
