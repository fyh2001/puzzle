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
    lastUnreadTotal: 0, // 上一次未读消息数量
    notifications: <NotificationListResp>{
      total: 0,
      records: [],
    },
  }),

  getters: {
    getNotificationList: (state) => state.notifications.records,
    getNotificationTotal: (state) => state.notifications.total,
    getNotificationUnreadTotal: (state) => state.unreadTotal,
    getLastUnreadTotal: (state) => state.lastUnreadTotal,
    getNotifiedStatus: (state) => state.notified,
  },

  actions: {
    // 不带查询条件地获取通知列表
    async fetchNotificationsDefault() {
      const userStore = useUserStore();

      const queryForm: NotificationReq = {
        userId: userStore.getUser.id,
        pagination: {
          page: 1,
          pageSize: 10,
        },
      };

      const {
        data: { code, data: notificationResp },
      } = await notificationRequest.list(queryForm);

      if (code === 200) {
        this.updateLastUnreadTotal();
        this.notifications.total = notificationResp.total;
        this.unreadTotal = notificationResp.records.filter(
          (item) => item.notificationUserStatusInfo.status === 0
        ).length;
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
      }
    },

    setNotified() {
      this.notified = true;
    },
    resetNotified() {
      this.notified = false;
    },
    updateLastUnreadTotal() {
      this.lastUnreadTotal = this.getNotificationUnreadTotal;
    },
    setLastUnreadTotal(value: number) {
      this.lastUnreadTotal = value;
    },
    reduceLastUnreadTotal() {
      this.lastUnreadTotal -= 1;
    },
  },
});
