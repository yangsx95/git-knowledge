package vo

type GetGitOrganizationsRequest struct {
	LoginInfo
	CredentialId string `url:"credential_id"`
}

type GitOrganization struct {
	Id        int64  `json:"id" xml:"id"`
	OrgId     string `json:"org_id" xml:"org_id"`
	AvatarUrl string `json:"avatar_url" xml:"avatar_url"`
}

type GetGitOrganizationsResponse []GitOrganization

type GetRepositoriesRequest struct {
	LoginInfo
	CredentialId   string `url:"credential_id"`
	OrganizationId string `url:"organization_id"`
}

type GitRepository struct {
	Name       string `json:"name" xml:"name"`
	FullName   string `json:"full_name" xml:"full_name"`
	HTMLUrl    string `json:"html_url" xml:"html_url"`
	CloneUrl   string `json:"clone_url" xml:"clone_url"`
	GitUrl     string `json:"git_url" xml:"git_url"`
	SSHUrl     string `json:"ssh_url" xml:"ssh_url"`
	Visibility string `json:"visibility" xml:"visibility"`
}

type GetRepositoriesResponse []GitRepository
