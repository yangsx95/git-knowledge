export class Request {
}

export class Response<T> {
  code: number
  msg: string
  detail: string;
  data: T
}
