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
