package db

import (
	"context"
	"github.com/joho/godotenv"
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

func InitResource(url, database string) (*Resource, error) {
	err := godotenv.Load(".env")
	if err != nil {
	}

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(url))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = mongoClient.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return &Resource{DB: mongoClient.Database(database)}, nil
}
