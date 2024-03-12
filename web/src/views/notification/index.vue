<script lang="ts" setup>
import { InfoRound } from "@vicons/material";
import { Raw, markRaw } from "vue";
import { useNotificationStore } from "@/store/notification";
import { notificationUserStatusRequest } from "api/notification";
import { useMessage } from "naive-ui";
import type { NotificationResp } from "@/types/notification";

const Message = useMessage();

const notificationStore = useNotificationStore();

const NotificationIcon: Record<string, Raw<any>> = {
  InfoRound: InfoRound,
};

const handleNotificationRead = async (data: NotificationResp) => {
  if (data.notificationUserStatusInfo.status === 1) return;

  const {
    data: { msg, code },
  } = await notificationUserStatusRequest.insert({ notificationId: data.id });

  if (code !== 200) {
    Message.error(msg);
  }

  notificationStore.fetchNotificationList({
    pagination: {
      page: 1,
      pageSize: 10,
    },
  });
};
console.log(notificationStore.getNotificationList);
</script>

<template>
  <div>
    <top-bar pt-4 title="通知" />

    <div class="p-4">
      <div
        v-for="(data, index) in notificationStore.getNotificationList"
        :key="index"
        @click="handleNotificationRead(data)"
      >
        <!-- 通知标注 -->
        <div class="flex justify-end h-2">
          <n-badge
            class="scale-150"
            dot
            processing
            :offset="[3.5, 4]"
            :show="data.notificationUserStatusInfo.status === 0"
          />
        </div>
        <!-- 通知内容 -->
        <div
          class="p-4 w-full bg-white shadow-sm rounded-xl mb-2 transition-all duration-300"
          :class="{
            'text-gray': data.notificationUserStatusInfo.status === 1,
          }"
        >
          <!-- 标题与时间 -->
          <div class="flex justify-between items-center">
            <div class="flex justify-between items-center gap-2">
              <n-icon
                :size="24"
                :component="NotificationIcon[data.notificationTypeInfo.icon]"
              />
              <div class="text-4.5 font-bold">
                {{ data.notificationTypeInfo?.name }}
              </div>
            </div>
            <div class="text-4 text-gray">
              {{ new Date(data.updatedAt).toLocaleDateString() }}
            </div>
          </div>
          <!-- 内容 -->
          <div class="mt-2 indent-4 text-4">{{ data.content }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>
