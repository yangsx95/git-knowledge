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
	UserApi  v1.UserApi
}

func initApi(app *App) *Api {
	api := Api{}

	api.LoginApi = v1.NewLoginApi(app.Dao.UserDao, app.Dao.ThirdPartOAuthDao)
	api.UserApi = v1.NewUserApi(app.Dao.UserDao)

	return &api
}

// initRouter 路由注册
func (a *App) initRouter() {
	e := a.echo
	api := a.Api

	// jwt认证中间件
	jm := middlewares.JWTMiddleware(os.Getenv("JWT_SECRET"))

	// v1 版本API
	groupV1 := e.Group("/api/v1")

	// 登录注册
	groupV1.POST("/registry", a.Handler(api.LoginApi.Registry))
	groupV1.POST("/login", a.Handler(api.LoginApi.Login))
	groupV1.GET("/oauth/authorize_url", a.Handler(api.LoginApi.GetOAuthAuthorizeUrl))
	groupV1.POST("/oauth/login", a.Handler(api.LoginApi.OAuthLogin))

	// user 用户
	groupV1.GET("/user", a.Handler(api.UserApi.GetUser), jm)
}
