package provider

import (
	"git-knowledge/api"
	"git-knowledge/api/request"
	"git-knowledge/dao"
	"git-knowledge/dao/model"
	"time"
)

type LoginService struct {
	userDao dao.UserDao
}

func NewLoginService(userDao dao.UserDao) api.LoginService {
	return &LoginService{
		userDao: userDao,
	}
}

func (l *LoginService) Registry(r request.RegistryRequest) error {
	user := model.User{
		Userid:    r.Userid,
		Password:  r.Password,
		Nickname:  r.Nickname,
		Email:     r.Email,
		Phone:     r.Phone,
		AvatarUrl: r.AvatarUrl,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
	err := l.userDao.InsertUser(user)
	return err
}
