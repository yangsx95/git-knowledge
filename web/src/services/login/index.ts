// @ts-ignore
/* eslint-disable */
import {API} from "@/services/login/typing";
import {request} from "umi";

/** 注册用户 POST /registry */
export async function register(body: API.RegisterParams, options?: { [key: string]: any }) {
  return request<API.RegisterResult>('/registry', {
    method: 'POST',
    data: body,
    ...(options || {}),
  });
}

/** 登录 */
export async function login(body: API.LoginParams, options?: { [key: string]: any }) {
  return request<API.LoginResult>('/login', {
    method: 'POST',
    data: body,
    ...(options || {}),
  });
}

/** 获取第三方oauth登录的url*/
export async function getOAuthAuthorizeUrl(type: string) {
  return request<API.GetOAuthAuthorizeUrlResult>('/oauth/url', {
    method: 'GET',
    params: {
      type: type
    }
  })
}

/** 三方oauth登录 */
export async function oauthLogin(body: API.OAuthLoginParams) {
  return request<API.OAuthLoginResult>('/oauth/login', {
    method: 'POST',
    data: body
  });
}
