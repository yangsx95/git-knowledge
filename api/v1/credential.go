package v1

import (
	"git-knowledge/api/v1/vo"
	"git-knowledge/dao"
	"git-knowledge/dao/model"
	"git-knowledge/result"
	"git-knowledge/util"
	"github.com/google/go-github/v39/github"
)

// CredentialApi Git凭据API
type CredentialApi interface {
	// GetGitOrganizations 获取Git组织
	GetGitOrganizations(request *vo.GetGitOrganizationsRequest) (*vo.GetGitOrganizationsResponse, error)

	// GetRepositories 获取指定git凭据下的所有仓库
	GetRepositories(request *vo.GetRepositoriesRequest) (*vo.GetRepositoriesResponse, error)
}

type credentialApiImpl struct {
	userDao dao.UserDao
}

func NewCredentialApi(userDao dao.UserDao) CredentialApi {
	return &credentialApiImpl{userDao: userDao}
}

func (c *credentialApiImpl) GetGitOrganizations(request *vo.GetGitOrganizationsRequest) (*vo.GetGitOrganizationsResponse, error) {
	err, user := c.userDao.FindUserByUserid(request.Userid)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, result.ErrorOf(result.CodeAuthErr)
	}

	cred := getCredentialById(request.CredentialId, user)
	if cred == nil {
		return nil, result.ErrorOf(result.CodeCredentialUseless)
	}

	resp := make(vo.GetGitOrganizationsResponse, 0)

	switch cred.ApiType {
	case "github":
		client, ctx := util.GetGithubClient(cred.Token)
		// 添加github用户自己
		resp = append(resp, vo.GitOrganization{
			Id:        user.Github.Id,
			OrgId:     user.Github.GithubId,
			AvatarUrl: "",
		})
		// 添加github用户所加入的组织
		orgs, _, err := client.Organizations.List(ctx, "", nil)
		if err != nil {
			return nil, err
		}
		for _, org := range orgs {
			resp = append(resp, vo.GitOrganization{
				Id:        *org.ID,
				OrgId:     *org.Login,
				AvatarUrl: *org.AvatarURL,
			})
		}
	}

	return &resp, nil
}

func getCredentialById(credentialId string, user *model.User) *model.Credential {
	cred := new(model.Credential)
	// 根据凭据id获取凭据
	switch credentialId {
	case "github":
		if user.Github.AccessToken == "" {
			return nil
		}
		cred.CredentialId = "github"
		cred.Name = "GitHub"
		cred.ApiType = "github"
		cred.Type = "access_token"
		cred.Token = user.Github.AccessToken
		return cred
	default:
	}
	return nil
}

func (c *credentialApiImpl) GetRepositories(request *vo.GetRepositoriesRequest) (*vo.GetRepositoriesResponse, error) {
	err, user := c.userDao.FindUserByUserid(request.Userid)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, result.ErrorOf(result.CodeAuthErr)
	}

	cred := getCredentialById(request.CredentialId, user)
	if cred == nil {
		return nil, result.ErrorOf(result.CodeCredentialUseless)
	}

	// 使用凭证读取仓库
	resp := make(vo.GetRepositoriesResponse, 0)
	switch cred.CredentialId {
	case "github":
		client, ctx := util.GetGithubClient(cred.Token)
		var repositories []*github.Repository
		if request.OrganizationId == user.Github.GithubId {
			repositories, _, err = client.Repositories.List(ctx, request.OrganizationId, nil)
		} else {
			repositories, _, err = client.Repositories.ListByOrg(ctx, request.OrganizationId, nil)
		}
		if err != nil {
			return nil, err
		}
		for _, r := range repositories {
			resp = append(resp, vo.GitRepository{
				Name:       *r.Name,
				FullName:   *r.FullName,
				HTMLUrl:    *r.HTMLURL,
				CloneUrl:   *r.CloneURL,
				GitUrl:     *r.GitURL,
				SSHUrl:     *r.SSHURL,
				Visibility: *r.Visibility,
			})
		}
	default:
	}
	return &resp, nil
}
