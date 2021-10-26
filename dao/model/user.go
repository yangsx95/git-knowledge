package model

import "time"

type User struct {
	Userid    string    `bson:"userid"`
	Password  string    `bson:"password"`
	Nickname  string    `bson:"nickname"`
	Email     string    `bson:"email"`
	Phone     string    `bson:"phone"`
	AvatarUrl string    `bson:"avatar_url"`
	CreatedAt time.Time `bson:"created_at"`
	UpdateAt  time.Time `bson:"update_at"`
	Github    Github    `bson:"github"`
}

type Github struct {
	Id          int64  `bson:"id"`
	AccessToken string `bson:"access_token"`
}
