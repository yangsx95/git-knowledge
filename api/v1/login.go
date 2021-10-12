package v1

import (
	"git-knowledge/dao"
	"git-knowledge/dao/model"
	"git-knowledge/result"
	"git-knowledge/util"
	"net/url"
	"os"
	"strconv"
	"time"
)

type LoginApi interface {

	// Registry 注册用户
	Registry(request *RegistryRequest) error

	// GetOAuthAuthorizeUrl 获取第三方oauth登录身份认证url，支持多种类型，比如github
	GetOAuthAuthorizeUrl(request *GetOAuthAuthorizeUrlRequest) *GetOAuthAuthorizeUrlResponse

	// OAuthLogin 用户授权成功后，调用此接口进行认证登录
	OAuthLogin(request *OAuthLoginRequest) error
}

type LoginApiImpl struct {
	userDao  dao.UserDao
	oAuthDao dao.OAuthDao
}

func NewLoginApi(userDao dao.UserDao, oAuthDao dao.OAuthDao) LoginApi {
	return &LoginApiImpl{userDao: userDao, oAuthDao: oAuthDao}
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

func (l *LoginApiImpl) GetOAuthAuthorizeUrl(request *GetOAuthAuthorizeUrlRequest) *GetOAuthAuthorizeUrlResponse {
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
	return &GetOAuthAuthorizeUrlResponse{
		Url: urlResult,
	}
}

func (l *LoginApiImpl) OAuthLogin(request *OAuthLoginRequest) error {
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
