// Package app 组件注册文件
package app

import (
	v1 "git-knowledge/api/v1"
	"git-knowledge/dao"
	"git-knowledge/middlewares"
	"os"
)

// Dao 应用程序组件容器，所有Dao组件都需要注册到该文件中
// 注意，要按照顺序依次注入
type Dao struct {
	UserDao           dao.UserDao
	ThirdPartOAuthDao dao.OAuthDao
}

func initDao(b *App) *Dao {
	d := Dao{}

	d.UserDao = dao.NewUserDao(b.db)
	d.ThirdPartOAuthDao = dao.NewThirdPartOAuthDao(b.db)

	return &d
}

// Api 组件注册对象
type Api struct {
	LoginApi v1.LoginApi
}

func initApi(b *App) *Api {
	a := Api{}

	a.LoginApi = v1.NewLoginApi(b.Dao.UserDao, b.Dao.ThirdPartOAuthDao)

	return &a
}

// initRouter 路由注册
func (a *App) initRouter() {
	r := a.echo
	api := a.Api
	r.POST("/api/registry", a.Handler(api.LoginApi.Registry))
	r.GET("/api/oauth/authorize_url", a.Handler(api.LoginApi.GetOAuthAuthorizeUrl))
	r.POST("/api/login/userid", a.Handler(api.LoginApi.LoginWithGitKnowledgeId))
	r.POST("/api/oauth/login", a.Handler(api.LoginApi.OAuthLogin))

	r.Group("/api/userinfo", middlewares.JWTMiddleware(os.Getenv("JWT_SECRET")))
}
