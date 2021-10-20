package vo

import "time"

type GetUserResponse struct {
	Userid    string    `json:"userid" xml:"userid"`
	Nickname  string    `json:"nickname" xml:"nickname"`
	Email     string    `json:"email" xml:"email"`
	Phone     string    `json:"phone" xml:"phone"`
	AvatarUrl string    `json:"avatar_url" xml:"avatar_url"`
	CreatedAt time.Time `json:"created_at" xml:"created_at"`
}

type GetUserRequest struct {
	LoginInfo
}
