// @ts-ignore
/* eslint-disable */
import {request} from 'umi';
import {API} from "@/services/user/typing";

/** 注册用户 POST /registry */
export async function register(body: API.RegisterParams, options?: { [key: string]: any }) {
  return request<API.RegisterResult>('/api/registry', {
    method: 'POST',
    data: body,
    ...(options || {}),
  });
}

/** 使用GitKnowledgeId登录 */
export async function loginWithGitKnowledgeId(body: API.LoginWithGitKnowledgeIdParams, options?: { [key: string]: any }) {
  return request<API.LoginWithGitKnowledgeIdResult>('/api/login/userid', {
    method: 'POST',
    data: body,
    ...(options || {}),
  });
}
