package dao

import (
	"git-knowledge/app/model"
	"git-knowledge/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDao interface {
	InsertUser(user model.User) error
}

type UserDaoImpl struct {
	resource   *db.Resource
	collection *mongo.Collection
}

func NewUserDao(resource *db.Resource) UserDao {
	return &UserDaoImpl{
		resource:   resource,
		collection: resource.DB.Collection("user"),
	}
}

func (u *UserDaoImpl) InsertUser(user model.User) error {
	context, cancel := initContext()
	defer cancel()
	_, err := u.collection.InsertOne(context, user)
	if err != nil {
		return err
	}
	return nil
}
