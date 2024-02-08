import { Pagination } from "@/types/pagination";
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
  id?: number;
  userId?: number;
  dimension?: number;
  type?: number;
  status?: number;
  durationRange?: [number, number];
  stepRange?: [number, number];
  dateRange?: [string, string];
  pagination?: Pagination;
  sorted?: string;
}

export interface RecordResp {
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

export interface RecordListResp {
  records: RecordModel[];
  total: number;
}
