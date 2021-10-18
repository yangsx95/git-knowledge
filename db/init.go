package db

import (
	"context"
	"fmt"
	"git-knowledge/util"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Resource struct {
	DB *mongo.Database
}

// Close 关闭连接
func (r *Resource) Close() {
}

func NewResource(host, port, database, username, password string) (*Resource, error) {
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "27017"
	}
	connectionUrl := ""
	if username != "" {
		connectionUrl = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	} else {
		connectionUrl = fmt.Sprintf("mongodb://%s:%s", host, port)
	}

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(connectionUrl))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = mongoClient.Connect(ctx)
	if err != nil {
		return nil, err
	}

	appDatabase := mongoClient.Database(database)
	initIndex(appDatabase)

	return &Resource{DB: appDatabase}, nil
}

func initIndex(client *mongo.Database) {
	ctx, cancel := util.GetContextWithTimeout60Second()
	defer cancel()
	one, err := client.Collection("user").Indexes().CreateMany(ctx, []mongo.IndexModel{{
		Keys: bson.M{
			"userid": 1,
		},
		Options: options.Index().SetUnique(true),
	}, {
		Keys: bson.M{
			"email": 1,
		},
		Options: options.Index().SetUnique(true),
	}})
	if err != nil {
		panic(err)
	}

	fmt.Println(one)
}
