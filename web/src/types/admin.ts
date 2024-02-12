export interface AuthReq {
  code: string;
}

export interface AuthResp {
  token: string;
}

export interface AuthSecretResp {
  secretUrl: string;
}
