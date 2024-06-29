import { request } from "@/api/index";
import type { HttpBaseResp } from "@/types/http";
import type {
  CreateRecordReq,
  RecordReq,
  RecordListResp,
  RecordBestSingleReq,
  RecordBestSingleListResp,
  RecordBestAverageReq,
  RecordBestAverageListResp,
  RecordBestStepReq,
  RecordBestStepListResp,
} from "@/types/record";

const baseURL = "/record";

export const recordRequest = {
  insert: (data: CreateRecordReq): Promise<HttpBaseResp<string>> =>
    request({
      url: `${baseURL}/insert`,
      method: "POST",
      data,
    }),

  listRecord: (data: RecordReq): Promise<HttpBaseResp<RecordListResp>> =>
    request({
      url: `${baseURL}/list-record`,
      method: "POST",
      data,
    }),

  listBestSingle: (
    data: RecordBestSingleReq
  ): Promise<HttpBaseResp<RecordBestSingleListResp>> =>
    request({
      url: `${baseURL}/list-best-single`,
      method: "POST",
      data,
    }),

  listBestAverage: (
    data: RecordBestAverageReq
  ): Promise<HttpBaseResp<RecordBestAverageListResp>> =>
    request({
      url: `${baseURL}/list-best-average`,
      method: "POST",
      data,
    }),

  listBestStep: (
    data: RecordBestStepReq
  ): Promise<HttpBaseResp<RecordBestStepListResp>> =>
    request({
      url: `${baseURL}/list-best-step`,
      method: "POST",
      data,
    }),
};
