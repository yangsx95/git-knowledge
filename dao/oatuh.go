package dao

import (
	"git-knowledge/dao/model"
	"git-knowledge/db"
	"git-knowledge/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type OAuthDao interface {
	Insert(auth model.OAuth) error
}

type oAuthDaoImpl struct {
	resource   *db.Resource
	collection *mongo.Collection
}

func (t *oAuthDaoImpl) Insert(auth model.OAuth) error {
	ctx, cancel := util.GetContextWithTimeout60Second()
	defer cancel()
	_, err := t.collection.InsertOne(ctx, auth)
	return err
}

func NewThirdPartOAuthDao(resource *db.Resource) OAuthDao {
	return &oAuthDaoImpl{
		resource:   resource,
		collection: resource.DB.Collection("oauth"),
	}
}
