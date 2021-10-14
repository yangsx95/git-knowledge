declare namespace API {
  type RegisterParams = {
    userid: string;
    password: string;
    nickname: string;
    email: string;
  };

  type RegisterResult = {
    code: number;
    msg: string;
    detail: string;
  };


}
