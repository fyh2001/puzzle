import { request } from "@/api/index";
import type {
  NotificationListResp,
  NotificationReq,
} from "@/types/notification";
import type { HttpBaseResp } from "@/types/http";

const baseURL = "/notification";

export const notificationRequest = {
  list: (data: NotificationReq): Promise<HttpBaseResp<NotificationListResp>> =>
    request({
      url: `${baseURL}/list`,
      method: "post",
      data,
    }),

  update: (data: NotificationReq): Promise<HttpBaseResp<string>> =>
    request({
      url: `${baseURL}/update`,
      method: "post",
      data,
    }),

  readAll: (data: NotificationReq): Promise<HttpBaseResp<string>> =>
    request({
      url: `${baseURL}/read-all`,
      method: "post",
      data,
    }),
};
