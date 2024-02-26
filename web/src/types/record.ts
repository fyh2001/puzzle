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
  durationFormat: string;
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
  id: string;
  userId: string;
  dimension: number;
  recordId: string;
  recordDuration: number;
  recordStep: number;
  recordBreakCount: number;
  ranked: number;
  createdAt: string;
  updatedAt: string;
}

export interface RecordBestSingleReq {
  id?: string;
  userId?: string;
  dimension?: number | null;
  recordId?: string;
  username?: string;
  nickname?: string;
  durationRange?: (number | string | null)[];
  stepRange?: (number | string | null)[];
  rankRange?: (number | string | null)[];
  dateRange?: [string, string];
  breakCountRange?: (number | string | null)[];
  pagination?: Pagination;
  sorted?: string;
  orderBy?: string;
  needUserInfo?: boolean;
  needRecordDetail?: boolean;
}

export interface RecordBestSingleResp {
  id: string;
  userId: string;
  dimension: number;
  recordId: string;
  recordDuration: number;
  recordStep: number;
  recordBreakCount: number;
  ranked: number;
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
  id: string;
  userId: string;
  dimension: number;
  type: number;
  recordIds: string;
  recordAverageDuration: number;
  recordBreakCount: number;
  ranked: number;
  createdAt: string;
  updatedAt: string;
}

export interface RecordBestAverageReq {
  id?: string;
  userId?: string;
  username?: string;
  nickname?: string;
  dimension?: number | null;
  type?: number | null;
  durationRange?: (number | string | null)[];
  rankRange?: (number | string | null)[];
  breakCountRange?: (number | string | null)[];
  dateRange?: [string, string];
  pagination?: Pagination;
  sorted?: string;
  orderBy?: string;
  needUserInfo?: boolean;
  needRecordDetail?: boolean;
}

export interface RecordBestAverageResp {
  id: string;
  userId: string;
  dimension: number;
  type: number;
  recordIds: string;
  recordAverageDuration: number;
  ranked: number;
  createdAt: string;
  updatedAt: string;
  userInfo: UserResp;
  recordDetail: RecordListResp;
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
