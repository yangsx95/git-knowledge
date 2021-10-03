package dao

import (
	"git-knowledge/dao/model"
	"git-knowledge/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		Id:        primitive.ObjectID{},
		Username:  "zhangsan",
		Password:  "123456",
		Name:      "张三",
		AvatarUrl: "我是图片地址",
		Phone:     "18362450836",
		CreatedAt: time.Time{},
	})
	if err != nil {
		panic(err)
	}
}
