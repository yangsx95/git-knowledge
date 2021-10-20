package v1

import (
	"git-knowledge/api/v1/vo"
	"git-knowledge/dao"
	"git-knowledge/result"
)

type UserApi interface {

	// GetUser 获取登录的用户信息
	GetUser(request *vo.GetUserRequest) (*vo.GetUserResponse, error)
}

type UserApiImpl struct {
	userDao dao.UserDao
}

func NewUserApi(userDao dao.UserDao) UserApi {
	return &UserApiImpl{userDao: userDao}
}

func (l *UserApiImpl) GetUser(request *vo.GetUserRequest) (*vo.GetUserResponse, error) {
	err, user := l.userDao.FindUserByUserid(request.Userid)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, result.ErrorOf(result.CodeAuthErr)
	}

	return &vo.GetUserResponse{
		Userid:    user.Userid,
		Nickname:  user.Nickname,
		Email:     user.Email,
		Phone:     user.Phone,
		AvatarUrl: user.AvatarUrl,
		CreatedAt: user.CreatedAt,
	}, nil
}
