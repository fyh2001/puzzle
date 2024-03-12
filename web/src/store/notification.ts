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
    lastUnreadTotal: 0, // 上一次未读消息数量
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
    getLastUnreadTotal: (state) => state.lastUnreadTotal,
    getNotifiedStatus: (state) => state.notified,
  },

  actions: {
    async fetchNotificationList(queryForm: NotificationReq) {
      const userStore = useUserStore();
      queryForm.userId = userStore.getUser.id;

      const {
        data: { code, data: notificationResp },
      } = await notificationRequest.list(queryForm);

      if (code === 200) {
        this.updateLastUnreadTotal();
        this.notifications.records = notificationResp.records;
        this.notifications.total = notificationResp.total;
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
