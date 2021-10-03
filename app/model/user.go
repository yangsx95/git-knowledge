package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	Name      string             `bson:"name"`
	AvatarUrl string             `bson:"avatar_url"`
	Phone     string             `bson:"phone"`
	CreatedAt time.Time          `bson:"created_at"`
}
