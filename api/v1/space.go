package v1

import (
	"git-knowledge/api/v1/vo"
	"git-knowledge/dao"
	"git-knowledge/dao/model"
)

type SpaceApi interface {
	// PostSpace 新建一个空间
	PostSpace(request *vo.PostSpaceRequest) error
}

type spaceApiImpl struct {
	spaceDao dao.SpaceDao
}

func NewSpaceApi(spaceDao dao.SpaceDao) SpaceApi {
	return &spaceApiImpl{spaceDao: spaceDao}
}

func (s *spaceApiImpl) PostSpace(request *vo.PostSpaceRequest) error {
	repos := make([]model.SpaceRepository, 0)
	for _, r := range *request.Repositories {
		repos = append(repos, model.SpaceRepository{
			RepositoryOwner: (*r.RepositoryId)[0],
			RepositoryId:    (*r.RepositoryId)[1],
			RepositoryName:  r.RepositoryName,
			CredentialId:    r.CredentialId,
		})
	}
	err := s.spaceDao.InsertOne(&model.Space{
		Name:         request.Name,
		Description:  request.Description,
		Owner:        request.Owner,
		Repositories: &repos,
	})
	return err
}
