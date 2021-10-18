import {CommonRequest, CommonResponse} from '../typing'

class RegisterRequest extends CommonRequest {
  userid: string;
  password: string;
  nickname: string;
  email: string;
}

class RegisterResponse extends CommonResponse {
}

class LoginWithGitKnowledgeIdRequest extends CommonRequest {
  userid: string;
  password: string;
}

class LoginWithGitKnowledgeIdResponse extends CommonResponse {
  token: string
}

declare namespace API {
  type RegisterParams = RegisterRequest;
  type RegisterResult = RegisterResponse;
  type LoginWithGitKnowledgeIdParams = LoginWithGitKnowledgeIdRequest;
  type LoginWithGitKnowledgeIdResult = LoginWithGitKnowledgeIdResponse;
}
