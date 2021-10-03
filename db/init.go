package db

import (
	"context"
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

func InitResource(url, database, username, password string) (*Resource, error) {
	client := options.Client()
	if username != "" {
		client.Auth.Username = username
	}
	if password != "" {
		client.Auth.Password = password
	}
	mongoClient, err := mongo.NewClient(client.ApplyURI(url))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = mongoClient.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &Resource{DB: mongoClient.Database(database)}, nil
}
