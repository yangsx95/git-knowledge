// Package app 组件注册文件
package app

import (
	v1 "git-knowledge/api/v1"
	"git-knowledge/dao"
)

// Dao 应用程序组件容器，所有Dao组件都需要注册到该文件中
// 注意，要按照顺序依次注入
type Dao struct {
	UserDao dao.UserDao
}

func initDao(b *App) *Dao {
	d := Dao{}

	d.UserDao = dao.NewUserDao(b.db)

	return &d
}

type Api struct {
	LoginApi v1.LoginApi
}

func initApi(b *App) *Api {
	a := Api{}

	a.LoginApi = v1.NewLoginApi(b.Dao.UserDao)

	return &a
}

func (a *App) initRouter() {
	r := a.echo
	api := a.Api

	r.POST("/registry", a.Handler(api.LoginApi.Registry))
}
