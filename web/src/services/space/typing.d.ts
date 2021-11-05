import type {Response} from '../typing';

class SpaceRequest {
  credential_id: string
  name: string
  type: string
}

class Space {
  description: string
  name: string
  owner: string
}

declare namespace API {
  type PostSpaceParam = SpaceRequest;
  type PostSpaceResult = Response<undefined>;
  type FindAllSpacesResult = Response<Space[]>;
}
