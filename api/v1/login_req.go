package v1

type RegistryRequest struct {
	Userid    string `json:"userid" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Nickname  string `json:"nickname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone"`
	AvatarUrl string `json:"avatar_url"`
}
