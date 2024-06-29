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
    notified: false, // 是否已经通知过
    unreadTotal: 0, // 未读消息数量
    notifications: <NotificationListResp>{
      total: 0,
      records: [],
    },
  }),

  getters: {
    getNotificationList: (state) => state.notifications.records,
    getNotificationTotal: (state) => state.notifications.total,
    getNotificationUnreadTotal: (state) => state.unreadTotal,
    getNotifiedStatus: (state) => state.notified,
  },

  actions: {
    // 不带查询条件地获取通知列表
    async fetchUnreadCount() {
      const userStore = useUserStore();

      const queryForm = {
        userId: userStore.getUser.id,
        readStatus: 1,
        onlyTotal: true,
      };

      const {
        data: { code, data: notificationResp },
      } = await notificationRequest.list(queryForm);

      if (code === 200) {
        this.unreadTotal = notificationResp.total;
      }
    },

    // 带查询条件地获取通知列表
    async fetchNotifications(queryForm: NotificationReq, type: string) {
      const userStore = useUserStore();
      queryForm.userId = userStore.getUser.id;

      const {
        data: { code, data: notificationResp },
      } = await notificationRequest.list(queryForm);

      if (code === 200) {
        if (type === "overwrite") {
          this.notifications.records = notificationResp.records;
        }
        if (type === "append") {
          this.notifications.records = [
            ...this.notifications.records,
            ...notificationResp.records,
          ];
        }
        this.notifications.total = notificationResp.total;
      }
    },

    setNotified() {
      this.notified = true;
    },
    resetNotified() {
      this.notified = false;
    },
  },
});
