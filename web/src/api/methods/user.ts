import { request } from "@/api/index";
import type {
  UserModel,
  UserListResp,
  UserloginReq,
  UserloginResp,
  UserReq,
} from "@/types/user";
import type { HttpBaseResp } from "@/types/http";

const baseURL = "/user";

export const userRequest = {
  register: (data: UserModel): Promise<HttpBaseResp<null>> =>
    request({
      url: `${baseURL}/register`,
      method: "post",
      data,
    }),

  login: (data: UserloginReq): Promise<HttpBaseResp<UserloginResp>> =>
    request({
      url: `${baseURL}/login`,
      method: "post",
      data,
    }),

  list: (data: UserReq): Promise<HttpBaseResp<UserListResp>> =>
    request({
      url: `${baseURL}/list`,
      method: "post",
      data,
    }),
};
