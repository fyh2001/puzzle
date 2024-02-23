export interface AuthReq {
  code: string;
}

export interface AuthResp {
  token: string;
  secretUrl: string;
}

export interface AuthSecretResp {
  secretUrl: string;
}

export interface AdminData<T> {
  list: T[];
  total: number;
}
