// @ts-ignore
/* eslint-disable */
import {API} from "@/services/user/typing";
import {request} from "umi";

/** 获取当前的登录信息 */
export async function getCurrentUser() {
  return request<API.GetCurrentUserResult>('/user', {
    method: 'GET',
  });
}

/** 获取登录用户的所有组织*/
export async function getOrganizations() {
  return request<API.GetOrganizationsResult>('/user/organizations', {
    method: 'GET',
  });
}

/** 获取登录用户的所有凭据*/
export async function getCredentials() {
  return request<API.GetCredentialsResult>('/user/credentials', {
    method: 'GET',
  });
}

/** 获取凭据下的组织 */
export async function getGitOrgs(cred_id: string) {
  return request<API.GetGitOrganizationResult>(`/credentials/${cred_id}/organizations`, {
    method: 'GET',
  });
}

/** 获取凭据下某个组织的所有仓库 */
export async function getGitRepos(cred_id: string, org_id: string) {
  return request<API.GetGitRepositoryResult>(`/credentials/${cred_id}/organizations/${org_id}/repositories`, {
    method: 'GET',
  });
}
