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
	UserDao  dao.UserDao
	SeqDao   dao.SeqDao
	SpaceDao dao.SpaceDao
}

func initDao(b *App) *Dao {
	d := Dao{}

	d.UserDao = dao.NewUserDao(b.db)
	d.SeqDao = dao.NewSeqDao(b.db)
	d.SpaceDao = dao.NewSpaceDao(b.db)

	return &d
}

// Api 组件注册对象
type Api struct {
	LoginApi      v1.LoginApi
	UserApi       v1.UserApi
	CredentialApi v1.CredentialApi
	SpaceApi      v1.SpaceApi
}

func initApi(app *App) *Api {
	api := Api{}

	api.LoginApi = v1.NewLoginApi(app.Dao.UserDao, app.Dao.SeqDao)
	api.UserApi = v1.NewUserApi(app.Dao.UserDao)
	api.CredentialApi = v1.NewCredentialApi(app.Dao.UserDao)
	api.SpaceApi = v1.NewSpaceApi(app.Dao.SpaceDao)
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
	groupV1.GET("/oauth/url", a.Handler(api.LoginApi.GetOAuthAuthorizeUrl))
	groupV1.POST("/oauth/login", a.Handler(api.LoginApi.OAuthLogin))

	// user 用户
	groupV1.GET("/user", a.Handler(api.UserApi.GetUser), jm)
	groupV1.GET("/user/organizations", a.Handler(api.UserApi.GetOrganizations), jm)
	groupV1.GET("/user/credentials", a.Handler(api.UserApi.GetCredentials), jm)

	// credential 凭证
	groupV1.GET("/credentials/:credential_id/organizations", a.Handler(api.CredentialApi.GetGitOrganizations), jm)
	groupV1.GET("/credentials/:credential_id/organizations/:organization_id/repositories", a.Handler(api.CredentialApi.GetRepositories), jm)

	// space 空间
	groupV1.POST("/space", a.Handler(api.SpaceApi.PostSpace), jm)
	groupV1.GET("/spaces", a.Handler(api.SpaceApi.ListAllByUserId), jm)
}
