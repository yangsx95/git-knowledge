package dao

import (
	"git-knowledge/dao/model"
	"git-knowledge/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDao interface {
	InsertUser(user model.User) error
}

type userDaoImpl struct {
	resource   *db.Resource
	collection *mongo.Collection
}

var userDaoInstance UserDao

func GetUserDaoInstance(resource *db.Resource) UserDao {
	if userDaoInstance == nil {
		userDaoInstance = &userDaoImpl{
			resource:   resource,
			collection: resource.DB.Collection("user"),
		}
	}
	return userDaoInstance
}

func (u *userDaoImpl) InsertUser(user model.User) error {
	context, cancel := initContext()
	defer cancel()
	_, err := u.collection.InsertOne(context, user)
	if err != nil {
		return err
	}
	return nil
}
