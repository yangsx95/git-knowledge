package dao

import (
	"git-knowledge/dao/model"
	"git-knowledge/db"
	"git-knowledge/util"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDao interface {
	InsertUser(user model.User) error
	FindUserByUserid(userid string) (error, *model.User)
	FindUserByEmail(email string) (error, *model.User)
	FindUserByGithubId(id int64) (error, *model.User)
	UpdateGithubInfo(userid string, githubInfo *model.Github) (int64, error)
}

type userDaoImpl struct {
	resource   *db.Resource
	collection *mongo.Collection
}

func NewUserDao(resource *db.Resource) UserDao {
	return &userDaoImpl{
		resource:   resource,
		collection: resource.DB.Collection("user"),
	}
}

func (u *userDaoImpl) InsertUser(user model.User) error {
	context, cancel := util.GetContextWithTimeout60Second()
	defer cancel()
	_, err := u.collection.InsertOne(context, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userDaoImpl) FindUserByUserid(userid string) (error, *model.User) {
	context, cancel := util.GetContextWithTimeout60Second()
	defer cancel()
	user := new(model.User)
	err := u.collection.FindOne(context, bson.M{"userid": userid}).Decode(user)
	if err != nil && err != mongo.ErrNoDocuments {
		return err, nil
	}
	if err != nil {
		return nil, nil
	}
	return nil, user
}

func (u *userDaoImpl) FindUserByEmail(email string) (error, *model.User) {
	context, cancel := util.GetContextWithTimeout60Second()
	defer cancel()
	user := new(model.User)
	err := u.collection.FindOne(context, bson.M{"email": email}).Decode(user)
	if err != nil && err != mongo.ErrNoDocuments {
		return err, nil
	}
	if err != nil {
		return nil, nil
	}
	return nil, user
}

func (u *userDaoImpl) FindUserByGithubId(id int64) (error, *model.User) {
	context, cancel := util.GetContextWithTimeout60Second()
	defer cancel()
	user := new(model.User)
	err := u.collection.FindOne(context, bson.M{"github.id": id}).Decode(user)
	if err != nil && err != mongo.ErrNoDocuments {
		return err, nil
	}
	if err != nil {
		return nil, nil
	}
	return nil, user
}

func (u *userDaoImpl) UpdateGithubInfo(userid string, githubInfo *model.Github) (int64, error) {
	context, cancel := util.GetContextWithTimeout60Second()
	defer cancel()
	one, err := u.collection.UpdateOne(context, bson.M{"userid": userid}, bson.M{"$set": bson.M{"github": githubInfo}})
	if err != nil {
		return 0, err
	}
	return one.ModifiedCount, nil
}
