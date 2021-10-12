package model

type OAuth struct {
	Channel     string `bson:"channel"`
	AccessToken string `bson:"access_token"`
	UserId      string `bson:"user_id"`
	AvatarURL   string `bson:"avatar_url"`
	Email       string `bson:"email"`
}
