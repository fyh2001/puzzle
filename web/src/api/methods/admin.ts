import { request } from "@/api/index";
import type { HttpBaseResp } from "@/types/http";
import type { AuthSecretResp, AuthReq, AuthResp } from "@/types/admin";
import type { UserReq, UserListResp } from "@/types/user";
import type {
  RecordReq,
  RecordListResp,
  RecordBestSingleReq,
  RecordBestSingleListResp,
  RecordBestAverageReq,
  RecordBestAverageListResp,
} from "@/types/record";
const baseURL = "/admin";

export const authRequest = {
  resetOtp: (): Promise<HttpBaseResp<AuthSecretResp>> =>
    request({
      url: `${baseURL}/authorization/reset-otp`,
      method: "post",
    }),

  authorization: (data: AuthReq): Promise<HttpBaseResp<AuthResp>> =>
    request({
      url: `${baseURL}/authorization/authorization`,
      method: "post",
      data,
    }),

  getUrl: (): Promise<HttpBaseResp<AuthSecretResp>> =>
    request({
      url: `${baseURL}/authorization/get-url`,
      method: "post",
    }),
};

export const userManageRequest = {
  list: (data: UserReq): Promise<HttpBaseResp<UserListResp>> =>
    request({
      url: `${baseURL}/user-manage/list`,
      method: "post",
      data,
    }),
  update: (data: UserReq): Promise<HttpBaseResp<string>> =>
    request({
      url: `${baseURL}/user-manage/update`,
      method: "post",
      data,
    }),
};

export const recordManageRequest = {
  list: (data: RecordReq): Promise<HttpBaseResp<RecordListResp>> =>
    request({
      url: `${baseURL}/record-manage/list`,
      method: "post",
      data,
    }),

  update: (data: RecordReq): Promise<HttpBaseResp<string>> =>
    request({
      url: `${baseURL}/record-manage/update`,
      method: "post",
      data,
    }),
};

export const recordBestSingleRequest = {
  list: (
    data: RecordBestSingleReq
  ): Promise<HttpBaseResp<RecordBestSingleListResp>> =>
    request({
      url: `${baseURL}/record-best-single-manage/list`,
      method: "post",
      data,
    }),
};

export const recordBestAverageRequest = {
  list: (
    data: RecordBestAverageReq
  ): Promise<HttpBaseResp<RecordBestAverageListResp>> =>
    request({
      url: `${baseURL}/record-best-average-manage/list`,
      method: "post",
      data,
    }),
};
