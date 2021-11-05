// @ts-ignore
/* eslint-disable */
import {API} from "@/services/space/typing";
import {request} from "umi";

/** 添加space */
export async function postSpace(body: API.PostSpaceParam) {
  return request<API.PostSpaceResult>('/space', {
    method: 'POST',
    data: body,
  });
}

export async function getAllSpaces() {
  return request<API.FindAllSpacesResult>('/spaces', {
    method: 'GET'
  });
}
