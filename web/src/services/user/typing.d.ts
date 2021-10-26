import type { Response} from '../typing';

class GetCurrentUserData {
  userid: string;
  nickname: string;
  email: string;
  phone: string;
  avatar_url: string;
  created_at: string;
}

declare namespace API {
  type GetCurrentUserResult = Response<GetCurrentUserData>;
  type CurrentUser = GetCurrentUserData;
}
