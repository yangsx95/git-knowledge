import type {Response} from '../typing';

class GetCurrentUserData {
  userid: string;
  nickname: string;
  email: string;
  phone: string;
  avatar_url: string;
  created_at: string;
}

class GetOrganizationData {
  org_id: string
  name: string;
  url: string;
  avatar_url: string;
}

class GetCredentialsData {
  credential_id: string
  name: string
  type: string
}

class GetGitOrganizationData {
  id: number
  org_id: string
  avatar_url: string
}

class GetGitRepositoryData {
  name: string
  full_name: string
  html_url: string
  clone_url: string
  git_url: string
  ssh_url: string
  visibility: string
}

declare namespace API {
  type GetCurrentUserResult = Response<GetCurrentUserData>;
  type CurrentUser = GetCurrentUserData;
  type GetOrganizationsResult = Response<GetOrganizationData[]>
  type GetCredentialsResult = Response<GetCredentialsData[]>
  type GetGitOrganizationResult = Response<GetGitOrganizationData[]>
  type GetGitRepositoryResult = Response<GetGitRepositoryData[]>
}
