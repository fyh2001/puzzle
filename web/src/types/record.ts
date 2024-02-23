import { Pagination } from "@/types/pagination";
import type { UserResp } from "./user";
// 个人记录
export interface RecordModel {
  id: string;
  userId: string;
  dimension: number;
  type: number;
  duration: number;
  step: number;
  status: number;
  scramble: string;
  solution: string;
  idx: string;
  createdAt: string;
  updatedAt: string;
}

export interface CreateRecordReq {
  dimension: number;
  type: number | undefined;
  duration: number;
  step: number;
  scramble: string;
  solution: string;
  idx: number | string;
}

export interface RecordReq {
  id?: string;
  userId?: string;
  dimension?: number | null;
  type?: number | null;
  duration?: number;
  step?: number;
  status?: number | null;
  scramble?: string;
  solution?: string;
  idx?: string;
  durationRange?: (number | string | null)[];
  stepRange?: (number | string | null)[];
  dateRange?: [string, string] | [Date, Date];
  pagination?: Pagination;
  sorted?: string;
  needUserInfo?: boolean;
}

export interface RecordResp {
  id: string;
  userId: string;
  userInfo: UserResp;
  dimension: number;
  type: number;
  duration: number;
  step: number;
  status: number;
  scramble: string;
  solution: string;
  idx: string;
  createdAt: string;
  updatedAt: string;
}

export interface RecordListResp {
  records: RecordResp[];
  total: number;
}

// 最佳单次
export interface RecordBestSingleModel {
  userId: string;
  dimension: number;
  recordId: string;
  recordDuration: number;
  recordStep: number;
  createdAt: string;
  updatedAt: string;
}

export interface RecordBestSingleReq {
  userId?: string;
  dimension?: number | null;
  recordId?: string;
  durationRange?: (number | string | null)[];
  stepRange?: (number | string | null)[];
  dateRange?: [string, string];
  pagination?: Pagination;
  sorted?: string;
  needUserInfo?: boolean;
  needRecordDetail?: boolean;
}

export interface RecordBestSingleResp {
  userId: string;
  dimension: number;
  recordId: string;
  recordDuration: number;
  recordStep: number;
  createdAt: string;
  updatedAt: string;
  userInfo: UserResp;
  recordDetail: RecordListResp;
}

export interface RecordBestSingleListResp {
  records: RecordBestSingleResp[];
  total: number;
}

// 最佳平均
export interface RecordBestAverageModel {
  userId: string;
  dimension: number;
  type: number;
  recordIds: string;
  recordAverageDuration: number;
  createdAt: string;
  updatedAt: string;
}

export interface RecordBestAverageReq {
  userId?: string;
  dimension?: number;
  type?: number;
  durationRange?: [number, number];
  dateRange?: [string, string];
  pagination?: Pagination;
  sorted?: string;
}

export interface RecordBestAverageResp {
  userId: string;
  userInfo: UserResp;
  dimension: number;
  type: number;
  recordIds: string;
  recordAverageDuration: number;
  createdAt: string;
  updatedAt: string;
}

export interface RecordBestAverageListResp {
  records: RecordBestAverageResp[];
  total: number;
}

// 最佳步数
export interface RecordBestStepModel {
  userId: string;
  dimension: number;
  recordId: string;
  recordStep: number;
  createdAt: string;
  updatedAt: string;
}

export interface RecordBestStepReq {
  userId?: number;
  dimension?: number;
  recordId?: number;
  stepRange?: [number, number];
  dateRange?: [string, string];
  pagination?: Pagination;
  sorted?: string;
}

export interface RecordBestStepResp {
  userId: string;
  userInfo: UserResp;
  dimension: number;
  recordId: string;
  recordStep: number;
  createdAt: string;
  updatedAt: string;
}

export interface RecordBestStepListResp {
  records: RecordBestStepResp[];
  total: number;
}
