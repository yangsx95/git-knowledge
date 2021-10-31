package vo

import "time"

type GetUserRequest struct {
	LoginInfo
}

type GetUserResponse struct {
	Userid    string    `json:"userid" xml:"userid"`
	Nickname  string    `json:"nickname" xml:"nickname"`
	Email     string    `json:"email" xml:"email"`
	Phone     string    `json:"phone" xml:"phone"`
	AvatarUrl string    `json:"avatar_url" xml:"avatar_url"`
	CreatedAt time.Time `json:"created_at" xml:"created_at"`
}

type GetOrganizationsRequest struct {
	LoginInfo
}

type Organization struct {
	OrgId     string `json:"org_id" xml:"org_id"`
	Name      string `json:"name" xml:"name"`
	Url       string `json:"url" xml:"url"`
	AvatarUrl string `json:"avatar_url" xml:"avatar_url"`
}

type GetOrganizationsResponse []*Organization

type GetCredentialsRequest struct {
	LoginInfo
}

type Credential struct {
	// 凭据id
	CredentialId string `json:"credential_id" xml:"credential_id"`
	// 凭据名称
	Name         string `json:"name" xml:"name"`
	// 凭据类型
	Type         string `json:"type" xml:"type"`
}

type GetCredentialsResponse []*Credential
