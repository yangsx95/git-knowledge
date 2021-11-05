package vo

type PostSpaceRequest struct {
	Name         string             `json:"name" xml:"name"`
	Description  string             `json:"description" xml:"description"`
	Owner        string             `json:"owner" xml:"owner"`
	Repositories *[]SpaceRepository `json:"repositories" xml:"repositories"`
}

type SpaceRepository struct {
	RepositoryId   *[]string `json:"repository_id" xml:"repository_id"`
	RepositoryName string    `json:"repository_name" xml:"repository_name"`
	CredentialId   string    `json:"credential_id" xml:"credential_id"`
}

type Space struct {
	Name         string             `json:"name" xml:"name"`
	Description  string             `json:"description" xml:"description"`
	Owner        string             `json:"owner" xml:"owner"`
	Repositories *[]SpaceRepository `json:"repositories" xml:"repositories"`
}

type ListAllByUserIdRequest struct {
	LoginInfo
}
