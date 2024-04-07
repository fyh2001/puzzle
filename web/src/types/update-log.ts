import { Pagination } from "@/types/pagination";

export interface UpdateLogModel {
  id: number;
  version: string;
  type: number;
  content: string;
  description: string;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export interface UpdateLogReq {
  id?: string;
  version?: string;
  type?: number | null;
  content?: string;
  description?: string;
  status?: number | null;
  dateRange?: [string, string];
  pagination?: Pagination;
  sorted?: string;
  orderBy?: string;
}

export interface UpdateLogResp {
  id: string;
  version: string;
  type: number;
  content: string | UpdateLogClassification;
  description: string;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export interface UpdateLogListResp {
  record: UpdateLogResp[];
  total: number;
}

export interface UpdateLogClassification {
  new: [];
  fix: [];
  optimize: [];
}
