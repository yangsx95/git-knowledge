import type { Response} from '../typing';
import {Request} from '../typing'

class RegisterRequest extends Request {
  userid: string;
  password: string;
  nickname: string;
  email: string;
}

class LoginRequest extends Request {
  userid: string;
  password: string;
}

class LoginData {
  token: string
}

class GetCurrentUserData {
  userid: string;
  nickname: string;
  email: string;
  phone: string;
  avatar_url: string;
  created_at: string;
}

declare namespace API {
  type RegisterParams = RegisterRequest;
  type RegisterResult = Response<undefined>;
  type LoginParams = LoginRequest;
  type LoginResult = Response<LoginData>;
  type GetCurrentUserResult = Response<GetCurrentUserData>;
  type CurrentUser = GetCurrentUserData;
}
