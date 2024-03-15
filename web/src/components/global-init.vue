<script lang="ts" setup>
import { onMounted } from "vue";
import { useNotificationStore } from "@/store/notification";
import { useUserStore } from "@/store/user";
import { useWebsocketStore } from "@/store/websocket";

const userStore = useUserStore();
const websocketStore = useWebsocketStore();
const notificationStore = useNotificationStore();

if (userStore.getToken) {
  websocketStore.connect();
}

onMounted(() => {
  notificationStore.fetchUnreadCount();
});
</script>

<template>
  <div></div>
</template>

<style scoped></style>
