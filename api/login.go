package api

import (
	"git-knowledge/api/request"
)

type LoginService interface {

	// Registry 注册用户
	Registry(request request.RegistryRequest) error
}
