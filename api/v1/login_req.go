package v1

type RegistryRequest struct {
	Userid    string `json:"userid" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Nickname  string `json:"nickname" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
	AvatarUrl string `json:"avatar_url"`
}
