import { defineStore } from "pinia";
import { useUserStore } from "./user";
import type {
  NotificationListResp,
  NotificationReq,
} from "@/types/notification";
import { notificationRequest } from "api/notification";

export const useNotificationStore = defineStore("notification", {
  // persist: true,

  state: () => ({
    notifications: <NotificationListResp>{
      total: 0,
      records: [],
    },
  }),

  getters: {
    getNotificationList: (state) => state.notifications.records,
    getNotificationTotal: (state) => state.notifications.total,
    getNotificationUnreadTotal: (state) => {
      return state.notifications.records.filter(
        (item) => item.notificationUserStatusInfo.status === 0
      ).length;
    },
  },

  actions: {
    async fetchNotificationList(queryForm: NotificationReq) {
      const userStore = useUserStore();
      queryForm.userId = userStore.getUser.id;

      const {
        data: { code, data: notificationResp },
      } = await notificationRequest.list(queryForm);

      if (code === 200) {
        this.notifications.records = notificationResp.records;
        this.notifications.total = notificationResp.total;
      }
    },
  },
});
