import { request } from "@/api/index";
import type {
  NotificationListResp,
  NotificationReq,
  NotificationUserStatusReq,
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
};

export const notificationUserStatusRequest = {
  insert: (data: NotificationUserStatusReq): Promise<HttpBaseResp<string>> =>
    request({
      url: `notification-user-status/insert`,
      method: "post",
      data,
    }),
};
