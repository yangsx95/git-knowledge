package dao

import (
	"git-knowledge/dao/model"
	"git-knowledge/db"
	"git-knowledge/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type SpaceDao interface {
	InsertOne(space *model.Space) error

	ListByUserId(userid string) (error, *[]model.Space)
}

type spaceDaoImpl struct {
	resource   *db.Resource
	collection *mongo.Collection
}

func NewSpaceDao(resource *db.Resource) SpaceDao {
	return &spaceDaoImpl{
		resource:   resource,
		collection: resource.DB.Collection("space"),
	}
}

func (s *spaceDaoImpl) InsertOne(space *model.Space) error {
	ctx, cancel := util.GetContextWithTimeout60Second()
	defer cancel()
	_, err := s.collection.InsertOne(ctx, space)
	return err
}

func (s *spaceDaoImpl) ListByUserId(userid string) (error, *[]model.Space) {
	panic("implement me")
}
