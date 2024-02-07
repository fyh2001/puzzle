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
