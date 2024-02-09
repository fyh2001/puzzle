import { request } from "@/api/index";

import type { HttpBaseResp } from "@/types/http";
import type { GetNewScrambleReq, ScrambleResp } from "@/types/scramble";

const baseURL = "/scramble";

export const scrambleRequest = {
  getNewScramble: (
    data: GetNewScrambleReq
  ): Promise<HttpBaseResp<ScrambleResp>> =>
    request({
      url: `${baseURL}/get-new-scramble`,
      method: "post",
      data,
    }),

  getUserScramble: (
    data: GetNewScrambleReq
  ): Promise<HttpBaseResp<ScrambleResp>> =>
    request({
      url: `${baseURL}/get-user-scramble`,
      method: "post",
      data,
    }),
};
