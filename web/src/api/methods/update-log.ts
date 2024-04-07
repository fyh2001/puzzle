import { request } from "@/api/index";
import type { HttpBaseResp } from "@/types/http";
import type {
  UpdateLogModel,
  UpdateLogReq,
  UpdateLogListResp,
} from "@/types/update-log";

const baseURL = "/update-log";

export const updateLogRequest = {
  insert: (data: UpdateLogModel): Promise<HttpBaseResp<string>> =>
    request({
      url: `${baseURL}/insert`,
      method: "POST",
      data,
    }),

  list: (data: UpdateLogReq): Promise<HttpBaseResp<UpdateLogListResp>> =>
    request({
      url: `${baseURL}/list`,
      method: "POST",
      data,
    }),
};
