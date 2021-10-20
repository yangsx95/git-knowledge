package vo

type RegistryRequest struct {
	Userid    string `json:"userid" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Nickname  string `json:"nickname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone"`
	AvatarUrl string `json:"avatar_url"`
}

type GetOAuthAuthorizeUrlRequest struct {
	Type string `query:"type" validate:"required"`
}

type OAuthLoginRequest struct {
	Type        string `json:"type" xml:"type"`
	Code        string `json:"code" xml:"code"`
	State       string `json:"state" xml:"state"`
	RedirectUrl string `json:"redirect_url" xml:"redirect_url"`
}

type GetOAuthAuthorizeUrlResponse struct {
	Url string `json:"url" xml:"url"`
}

type OAuthLoginResponse struct {
}

type LoginRequest struct {
	Userid   string `json:"userid" xml:"userid" validate:"required"`
	Password string `json:"password" xml:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token" xml:"token"`
}
