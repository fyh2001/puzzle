import { request } from "@/api/index";
import type { HttpBaseResp } from "@/types/http";
import type { AuthSecretResp, AuthReq, AuthResp } from "@/types/admin";
const baseURL = "/admin";

export const authRequest = {
  resetOtp: (): Promise<HttpBaseResp<AuthSecretResp>> =>
    request({
      url: `${baseURL}/authorization/reset-otp`,
      method: "POST",
    }),

  authorization: (data: AuthReq): Promise<HttpBaseResp<AuthResp>> =>
    request({
      url: `${baseURL}/authorization/authorization`,
      method: "POST",
      data,
    }),

  getUrl: (): Promise<HttpBaseResp<AuthSecretResp>> =>
    request({
      url: `${baseURL}/authorization/get-url`,
      method: "POST",
    }),
};
