<script lang="ts" setup>
import { useMessage } from "naive-ui";
import { useIntervalAsync } from "@/utils/useIntervalAsync";
import { useNotificationStore } from "@/store/notification";

const Message = useMessage();
const notificationStore = useNotificationStore();

const polling = async () => {
  fetchNotifications();
};

const fetchNotifications = async () => {
  await notificationStore.fetchNotificationList({
    pagination: {
      page: 1,
      pageSize: 10,
    },
  });

  // 有未读消息
  if (notificationStore.getNotificationUnreadTotal > 0) {
    //首次加载
    if (!notificationStore.getNotifiedStatus) {
      notificationStore.setNotified();
      return Message.info(
        `您有 ${notificationStore.getNotificationUnreadTotal} 条未读消息`
      );
    }

    // 非首次加载
    if (
      notificationStore.getNotifiedStatus &&
      notificationStore.getNotificationUnreadTotal -
        notificationStore.getLastUnreadTotal >
        0
    ) {
      return Message.info(
        `您有 ${
          notificationStore.getNotificationUnreadTotal -
          notificationStore.getLastUnreadTotal
        } 条新的未读消息`
      );
    }
  }
};

useIntervalAsync(polling, 10000);
</script>

<template>
  <div></div>
</template>

<style scoped></style>
