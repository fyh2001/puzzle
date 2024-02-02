<script lang="ts" setup>
import {
  VideogameAssetOutlined,
  InsertChartOutlinedRound,
  PersonOutlineRound,
} from "@vicons/material";
import Tabbar from "@/components/tabbar/index.vue";
import type { TabbarItem } from "@/components/tabbar/types";

const tabsContent: TabbarItem[] = [
  {
    label: "练习",
    icon: VideogameAssetOutlined,
    path: "/",
  },
  {
    label: "记录",
    icon: InsertChartOutlinedRound,
    path: "/record",
  },
  {
    label: "我的",
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
