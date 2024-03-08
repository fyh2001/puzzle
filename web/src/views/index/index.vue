<script lang="ts" setup>
import {
  VideogameAssetOutlined,
  InsertChartOutlinedRound,
  PersonOutlineRound,
} from "@vicons/material";
import Tabbar from "@/components/tabbar.vue";
import type { TabbarItem } from "@/components/tabbar.vue";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const tabsContent: TabbarItem[] = [
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
];
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
