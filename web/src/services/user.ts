// @ts-ignore
/* eslint-disable */
import {request} from 'umi';

/** 注册用户 POST /registry */
export async function register(body: API.RegisterParams, options?: { [key: string]: any }) {
  return request<API.RegisterResult>('/api/registry', {
    method: 'POST',
    data: body,
    ...(options || {}),
  });
}
