// @ts-ignore
/* eslint-disable */
import {API} from "@/services/user/typing";
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

/** 获取当前的登录信息 */
export async function getCurrentUser() {
  return request<API.GetCurrentUserResult>('/user', {
    method: 'GET',
  });
}
