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

class GetOAuthAuthorizeUrlData {
  url: string
}

class OAuthLoginRequest extends Request {
  code: string;
  type: string;
  state: string;
  redirect_url: string;
}

class OAuthLoginData {
  token: string;
}

declare namespace API {
  type RegisterParams = RegisterRequest;
  type RegisterResult = Response<undefined>;
  type LoginParams = LoginRequest;
  type LoginResult = Response<LoginData>;
  type GetOAuthAuthorizeUrlResult = Response<GetOAuthAuthorizeUrlData>;
  type OAuthLoginParams = OAuthLoginRequest;
  type OAuthLoginResult = Response<OAuthLoginData>;
}
