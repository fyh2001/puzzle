<script lang="ts" setup>
import { computed } from "vue";
import {
  VideogameAssetOutlined,
  InsertChartOutlinedRound,
  PersonOutlineRound,
} from "@vicons/material";
import Tabbar from "@/components/tabbar.vue";
import { useI18n } from "vue-i18n";
import { useMessage } from "naive-ui";
import { useUserStore } from "@/store/user";
import { useNotificationStore } from "@/store/notification";
import type { TabbarItem } from "@/components/tabbar.vue";

const { t } = useI18n();
const Message = useMessage();

const tabsContent = computed<TabbarItem[]>(() => [
  {
    label: t("tabbar.practice"),
    icon: VideogameAssetOutlined,
    path: "/",
  },
  {
    label: t("tabbar.record"),
    icon: InsertChartOutlinedRound,
    path: "/record",
  },
  {
    label: t("tabbar.user"),
    icon: PersonOutlineRound,
    path: "/user",
  },
]);

const userStore = useUserStore();
const notificationStore = useNotificationStore();

if (userStore.getToken) {
  notificationStore
    .fetchNotificationList({
      pagination: {
        page: 1,
        pageSize: 10,
      },
      userId: userStore.getUser.id,
    })
    .then(() => {
      if (notificationStore.getNotificationUnreadTotal > 0)
        Message.info(
          `您有 ${notificationStore.getNotificationUnreadTotal} 条未读消息`
        );
    });
}
</script>

<template>
  <div>
    <div class="h-100vh overflow-auto">
      <router-view v-slot="{ Component, route }">
        <transition
          :name="route.meta.transition as string || 'fade'"
          mode="default"
        >
          <component :is="Component" class="w-screen" />
        </transition>
      </router-view>
    </div>
    <div class="fixed bottom-0 w-full p-3 transition-all duration-300 z-100">
      <tabbar :tabsContent="tabsContent" />
    </div>
  </div>
</template>
