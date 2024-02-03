import { request } from "@/api/index";
import type { UserModel, UserloginReq, UserloginResp } from "@/types/user";
import type { HttpBaseResp } from "@/types/http";

const baseURL = "/user";

export const userRequest = {
  register: (data: UserModel): Promise<HttpBaseResp<null>> =>
    request({
      url: `${baseURL}/register`,
      method: "POST",
      data,
    }),

  login: (data: UserloginReq): Promise<HttpBaseResp<UserloginResp>> =>
    request({
      url: `${baseURL}/login`,
      method: "POST",
      data,
    }),
};
