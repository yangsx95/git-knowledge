package dao

import (
	"git-knowledge/dao/model"
	"git-knowledge/db"
	"git-knowledge/util"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SpaceDao interface {
	InsertOne(space *model.Space) error

	ListByUserId(userid string) (*[]model.Space, error)
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

func (s *spaceDaoImpl) ListByUserId(userid string) (*[]model.Space, error) {
	ctx, cancel := util.GetContextWithTimeout60Second()
	defer cancel()
	cur, err := s.collection.Find(ctx, bson.M{"owner": userid})
	if err != nil {
		return nil, err
	}
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}
	if err != nil {
		return nil, nil
	}
	spaces := make([]model.Space, 0)
	for cur.Next(ctx) {
		var b model.Space
		if err = cur.Decode(&b); err != nil {
			return nil, err
		}
		spaces = append(spaces, b)
	}
	return &spaces, nil
}
