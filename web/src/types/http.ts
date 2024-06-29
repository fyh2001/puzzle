export interface HttpResp<T> {
  code: number;
  msg: string;
  data: T;
}

export interface HttpBaseResp<T> {
  status: number;
  data: HttpResp<T>;
}
