package v1

import (
	"git-knowledge/api/v1/vo"
	"git-knowledge/dao"
	"git-knowledge/result"
)

type UserApi interface {

	// GetUser 获取登录的用户信息
	GetUser(request *vo.GetUserRequest) (*vo.GetUserResponse, error)

	// GetOrganizations 获取登录用户的所有组织
	GetOrganizations(req *vo.GetOrganizationsRequest) (*vo.GetOrganizationsResponse, error)

	// GetCredentials 获取登录用户的配置的所有API凭据
	GetCredentials(req *vo.GetCredentialsRequest) (*vo.GetCredentialsResponse, error)
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

func (l *UserApiImpl) GetOrganizations(request *vo.GetOrganizationsRequest) (*vo.GetOrganizationsResponse, error) {
	err, user := l.userDao.FindUserByUserid(request.Userid)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, result.ErrorOf(result.CodeAuthErr)
	}
	resp := new(vo.GetOrganizationsResponse)
	*resp = append(*resp, &vo.Organization{
		OrgId:     user.Userid,
		Name:      user.Nickname,
		Url:       "",
		AvatarUrl: user.AvatarUrl,
	})
	return resp, nil
}

func (l *UserApiImpl) GetCredentials(request *vo.GetCredentialsRequest) (*vo.GetCredentialsResponse, error) {
	err, user := l.userDao.FindUserByUserid(request.Userid)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, result.ErrorOf(result.CodeAuthErr)
	}
	resp := new(vo.GetCredentialsResponse)

	// 登录使用的凭据（github、gitlab、gitee）
	*resp = append(*resp, &vo.Credential{
		CredentialId: "github",
		Name:         "Github",
		Type:         "token",
	})
	return resp, nil
}
