package dao

import (
	"git-knowledge/dao/model"
	"git-knowledge/db"
	"git-knowledge/util"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
)

type SeqDao interface {
	GenUserId() (string, error)
}

type seqDaoImpl struct {
	resource   *db.Resource
	collection *mongo.Collection
}

func NewSeqDao(resource *db.Resource) SeqDao {
	return &seqDaoImpl{
		resource:   resource,
		collection: resource.DB.Collection("seq"),
	}
}

func (s *seqDaoImpl) GenUserId() (string, error) {
	ctx, cancel := util.GetContextWithTimeout60Second()
	defer cancel()
	seq := model.Seq{}
	rt := s.collection.FindOneAndUpdate(ctx,
		bson.M{"seq_name": "userid_gen"},
		bson.M{"$inc": bson.M{"seq_val": 1}},
	)
	err := rt.Decode(&seq)
	if err == mongo.ErrNoDocuments {
		_, err := s.collection.InsertOne(ctx, bson.M{"seq_name": "userid_gen", "seq_val": 1})
		if err != nil {
			return "", err
		}
		return s.GenUserId()
	} else if err != nil {
		return "", err
	}
	return "git_knowledge_" + strconv.Itoa(seq.SeqVal), nil
}
