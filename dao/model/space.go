package model

type Space struct {
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	Owner        string             `bson:"owner"`
	Repositories *[]SpaceRepository `bson:"repositories"`
}

type SpaceRepository struct {
	RepositoryOwner string `bson:"repository_owner"`
	RepositoryId    string `bson:"repository_id"`
	RepositoryName  string `bson:"repository_name"`
	CredentialId    string `bson:"credential_id"`
}
