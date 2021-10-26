package dao

import (
	"fmt"
	"git-knowledge/dao/model"
	"git-knowledge/db"
	"testing"
	"time"
)

func InitUserDao() UserDao {
	resource, err := db.NewResource("127.0.0.1", "27017", "app", "root", "root123")
	if err != nil {
		panic(err)
	}
	return NewUserDao(resource)
}

func TestUserDaoImpl_InsertUser(t *testing.T) {
	dao := InitUserDao()
	err := dao.InsertUser(model.User{
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

func TestUserDaoImpl_UpdateUserGithubAccessToken(t *testing.T) {
	dao := InitUserDao()
	count, err := dao.UpdateUserGithubAccessToken("123456", "123456_test_access_token")
	if err != nil {
		panic(err)
	}
	fmt.Printf("更新条数：%v", count)
}
