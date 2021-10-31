package model

type Credential struct {
	CredentialId string `bson:"credential_id"`
	Name         string `bson:"name"`
	Type         string `bson:"type"`
	ApiType      string `bson:"api_type"`
	Token        string `bson:"token"`
}
