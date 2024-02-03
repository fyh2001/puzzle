import { Pagination } from "@/types/pagination";

export interface UserModel {
  id?: string;
  username?: string;
  password?: string;
  nickname?: string;
  email?: string;
  phone?: string;
  avatar?: string;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
}

export interface UserReq {
  id?: string;
  username?: string;
  password?: string;
  nickname?: string;
  email?: string;
  phone?: string;
  status?: number;
  pagination?: Pagination;
}

export interface UserResp {
  id: string;
  username: string;
  nickname: string;
  email: string;
  phone: string;
  avatar: string;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export interface UserListResp {
  records: UserResp[];
  total: number;
}

export interface UserloginReq {
  username: string;
  password: string;
}

export interface UserloginResp {
  token: string;
  user: UserResp;
}

export interface UserRegisterReq {
  username: string;
  password: string;
  nickname: string;
}
