import type {Response} from '../typing';

class SpaceRequest {
  credential_id: string
  name: string
  type: string
}

declare namespace API {
  type PostSpaceParam = SpaceRequest;
  type PostSpaceResult = Response<undefined>;
}
