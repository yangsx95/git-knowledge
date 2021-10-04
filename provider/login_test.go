package provider

import (
	"git-knowledge/api"
	"git-knowledge/api/request"
	"git-knowledge/dao"
	"git-knowledge/db"
	"testing"
)

var instance api.LoginService

func init() {
	resource, err := db.NewResource("127.0.0.1", "27017", "test", "root", "root123")
	if err != nil {
		panic(err)
	}
	userDao := dao.NewUserDao(resource)
	instance = NewLoginService(userDao)
}

func TestLoginService_Registry(t *testing.T) {
	req := request.RegistryRequest{
		Userid:    "xf616510229",
		Password:  "123456",
		Nickname:  "张三",
		Email:     "616510229@qq.com",
		Phone:     "18362123334",
		AvatarUrl: "",
	}
	err := instance.Registry(req)
	if err != nil {
		panic(err)
	}
}
