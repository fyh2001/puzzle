<script lang="ts" setup>
import GlobalProvider from "@/components/global-provider.vue";
import { onMounted } from "vue";
import { useThemeStore } from "@/store/theme";
import { useNotificationStore } from "./store/notification";

const themeStore = useThemeStore();
const notificationStore = useNotificationStore();

onMounted(() => {
  notificationStore.fetchUnreadCount();
});
</script>

<template>
  <global-provider>
    <router-view v-slot="{ Component }">
      <transition :name="'fade'" mode="out-in">
        <component
          :is="Component"
          :class="{ 'brightness-60': themeStore.darkMode }"
        />
      </transition>
    </router-view>
  </global-provider>
</template>

<style></style>
