package dao

import (
	"git-knowledge/dao/model"
	"git-knowledge/db"
	"testing"
	"time"
)

func TestUserDaoImpl_InsertUser(t *testing.T) {

	resource, err := db.InitResource("127.0.0.1", "27017", "test", "root", "root123")
	if err != nil {
		panic(err)
	}

	dao := GetUserDaoInstance(resource)
	err = dao.InsertUser(model.User{
		Userid:    "root",
		Password:  "root123",
		Nickname:  "管理员",
		Email:     "root@qq.com",
		Phone:     "18878092222",
		AvatarUrl: "",
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
		Github:    model.Github{AccessToken: "123"},
	})
	if err != nil {
		panic(err)
	}
}
