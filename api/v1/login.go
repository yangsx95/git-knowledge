package v1

import (
	"git-knowledge/api/v1/vo"
	"git-knowledge/dao"
	"git-knowledge/dao/model"
	"git-knowledge/middlewares"
	"git-knowledge/result"
	"git-knowledge/util"
	"github.com/golang-jwt/jwt"
	"net/url"
	"os"
	"strconv"
	"time"
)

type LoginApi interface {

	// Registry 注册用户
	Registry(request *vo.RegistryRequest) error

	// Login 登录
	Login(request *vo.LoginRequest) (*vo.LoginResponse, error)

	// GetOAuthAuthorizeUrl 获取第三方oauth登录身份认证url，支持多种类型，比如github
	GetOAuthAuthorizeUrl(request *vo.GetOAuthAuthorizeUrlRequest) *vo.GetOAuthAuthorizeUrlResponse

	// OAuthLogin 用户授权成功后，调用此接口进行认证登录
	OAuthLogin(request *vo.OAuthLoginRequest) error
}

type LoginApiImpl struct {
	userDao  dao.UserDao
	oAuthDao dao.OAuthDao
}

func NewLoginApi(userDao dao.UserDao, oAuthDao dao.OAuthDao) LoginApi {
	return &LoginApiImpl{userDao: userDao, oAuthDao: oAuthDao}
}

func (l *LoginApiImpl) Registry(request *vo.RegistryRequest) error {
	// 判断邮箱以及用户id是否存在
	err, user := l.userDao.FindUserByUserid(request.Userid)
	if err != nil {
		return err
	}
	if user != nil {
		return result.ErrorOf(result.CodeRegisterUserIdAlreadyExists)
	}
	err, user = l.userDao.FindUserByEmail(request.Email)
	if err != nil {
		return err
	}
	if user != nil {
		return result.ErrorOf(result.CodeRegisterEmailAlreadyExists)
	}

	err = l.userDao.InsertUser(model.User{
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

func (l *LoginApiImpl) Login(request *vo.LoginRequest) (*vo.LoginResponse, error) {
	var err error
	var user *model.User
	if util.IsEmailAddr(request.Userid) {
		err, user = l.userDao.FindUserByEmail(request.Userid)
	} else {
		err, user = l.userDao.FindUserByUserid(request.Userid)
	}
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, result.ErrorOf(result.CodeUserNotExists)
	}
	if user.Password != request.Password {
		return nil, result.ErrorOf(result.CodeWrongPassword)
	}
	// 登录成功，生成jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middlewares.JWTClaims{
		Userid: user.Userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // 有效期为72小时
		},
	})
	// 生成token字符串
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	return &vo.LoginResponse{
		Token: t,
	}, nil
}

func (l *LoginApiImpl) GetOAuthAuthorizeUrl(request *vo.GetOAuthAuthorizeUrlRequest) *vo.GetOAuthAuthorizeUrlResponse {
	urlResult := ""
	switch request.Type {
	case "github":
		// 获取用户身份认证url
		// https://docs.github.com/cn/developers/apps/building-oauth-apps/authorizing-oauth-apps
		u, _ := url.Parse("https://github.com/login/oauth/authorize")
		query := u.Query()
		query.Add("client_id", os.Getenv("GITHUB_CLIENT_ID"))
		query.Add("redirect_uri", os.Getenv("GITHUB_REDIRECT_URI"))
		query.Add("scope", os.Getenv("GITHUB_SCOPE"))
		query.Add("state", util.RandStr(6))
		// 解析RawQuery并返回"值，您得到的只是URL查询值的副本，而不是"实时引用"，
		// 因此修改该副本不会对原始查询产生任何影响。
		// 为了修改原始查询，您必须分配给原始RawQuery
		u.RawQuery = query.Encode()
		urlResult = u.String()
	}
	return &vo.GetOAuthAuthorizeUrlResponse{
		Url: urlResult,
	}
}

func (l *LoginApiImpl) OAuthLogin(request *vo.OAuthLoginRequest) error {
	switch request.Type {
	case "github":
		// 获取accessToken
		resp, err := util.GetGithubAccessToken(os.Getenv("GITHUB_CLIENT_ID"), os.Getenv("GITHUB_SECRET"), request.Code, request.RedirectUrl)
		if err != nil {
			return err
		}
		if resp.Error != "" {
			return result.ErrorOfWithDetail(result.CodeGithubAuthFail, resp.ErrorDescription)
		}
		// 根据token获取用户信息
		client, ctx := util.GetGithubClient(resp.AccessToken)
		user, _, err := client.Users.Get(ctx, "")
		if err != nil {
			return err
		}
		// 存储token以及第三方(github)登录信息
		err = l.oAuthDao.Insert(model.OAuth{
			Channel:     "github",
			AccessToken: resp.AccessToken,
			UserId:      strconv.FormatInt(*user.ID, 10),
			AvatarURL:   *user.AvatarURL,
			Email:       *user.Email,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
