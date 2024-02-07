import { request } from "@/api/index";
import type { HttpBaseResp } from "@/types/http";
import type { CreateRecordReq } from "@/types/record";

const baseURL = "/record";

export const recordRequest = {
  insert: (data: CreateRecordReq): Promise<HttpBaseResp<string>> =>
    request({
      url: `${baseURL}/insert`,
      method: "POST",
      data,
    }),
};
