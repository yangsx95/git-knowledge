package v1

import (
	"git-knowledge/dao"
	"git-knowledge/dao/model"
	"time"
)

type LoginApi interface {

	// Registry 注册用户
	Registry(request *RegistryRequest) error
}

type LoginApiImpl struct {
	userDao dao.UserDao
}

func NewLoginApi(userDao dao.UserDao) LoginApi {
	return &LoginApiImpl{userDao: userDao}
}

func (l *LoginApiImpl) Registry(request *RegistryRequest) error {
	err := l.userDao.InsertUser(model.User{
		Userid:    request.Userid,
		Password:  request.Password,
		Nickname:  request.Nickname,
		Email:     request.Email,
		Phone:     request.Phone,
		AvatarUrl: request.AvatarUrl,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	})
	return err
}
