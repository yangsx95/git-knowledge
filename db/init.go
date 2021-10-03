package db

import (
	"context"
	"fmt"
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

func InitResource(host, port, database, username, password string) (*Resource, error) {
	if host == "" {
		host = "127.0.0.1"
	}
	if port == "" {
		port = "27017"
	}
	connectionUrl := ""
	if username != "" {
		connectionUrl = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", username, password, host, port, database)
	} else {
		connectionUrl = fmt.Sprintf("mongodb://%s:%s/%s", host, port, database)
	}

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(connectionUrl))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = mongoClient.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &Resource{DB: mongoClient.Database(database)}, nil
}
