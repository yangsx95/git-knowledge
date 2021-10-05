// Package app 组件注册文件
package app

import (
	"fmt"
	v1 "git-knowledge/api/v1"
	"git-knowledge/dao"
	"reflect"
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

type Api struct {
	LoginApi v1.LoginApi
}

func initApi(b *BootStrap) *Api {
	a := Api{}
	a.LoginApi = v1.NewLoginApi(b.Dao.UserDao)
	return &a
}

func initRouter(b *BootStrap) {
	//r := b.engine.RouterGroup
	//l := b.Api.LoginApi
	//r.POST("/registry", func(context *gin.Context) {
	//	request := v1.RegistryRequest{}
	//	rr := b.Api.LoginApi.Registry
	//	context.Bind(&request)
	//	err := l.Registry(request)
	//})
}

func Handler(apiMethod interface{}) {
	mT := reflect.TypeOf(apiMethod)
	mV := reflect.ValueOf(apiMethod)

	// 参数大于1个
	pLen := mT.NumIn()
	if pLen > 1 {
		panic("API 方法不符合规范，最多只能拥有一个参数" + mT.Name())
	}

	// 参数等于1个，需要生成参数对象并调用
	pVs := make([]reflect.Value, 0)
	if pLen != 0 {
		pT := mT.In(0).Elem()
		pV := reflect.New(pT)
		fmt.Println(pV)
		pVs = append(pVs, pV)
	}

	call := mV.Call(pVs)
	fmt.Println(call)
}
