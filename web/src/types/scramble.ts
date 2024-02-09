export interface ScrambleModel {
  scrambleMap: number[];
  scrambleStr: string;
  scrambleIdx: number;
}

export interface GetNewScrambleReq {
  dimension: number;
}

export interface ScrambleResp {
  id: string;
  dimension: number;
  idx: number;
  scramble: string;
  status: number;
  createdAt: string;
  updatedAt: string;
}

export interface ScrambleListResp {
  records: ScrambleResp[];
  total: number;
}
