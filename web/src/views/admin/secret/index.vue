<script setup lang="ts">
import { onMounted } from "vue";
import { useAdminStore } from "@/store/admin";
import { useDialog } from "naive-ui";

const Dialog = useDialog();

const adminStore = useAdminStore();

const getSecretUrl = async () => {
  adminStore.fetchSecretUrl();
};

const handleSecretUrlDetail = () => {
  Dialog.info({
    title: "秘钥详情",
    content: adminStore.getSecretUrl,
    style: {
      width: "30%",
    },
  });
};

const handleResetSecretUrl = () => {
  Dialog.warning({
    title: "重置秘钥",
    content: "确定要重置秘钥吗？",
    positiveText: "确定",
    onPositiveClick: () => {
      adminStore.resetOtp();
    },
  });
};

onMounted(() => {
  getSecretUrl();

  console.log(adminStore.getSecretUrl);
});
</script>

<template>
  <div
    class="absolute left-1/2 top-1/2 -translate-x-1/2 -translate-y-2/3 flex flex-col gap-4"
  >
    <div class="p-4 bg-white rounded-xl">
      <n-qr-code :value="adminStore.getSecretUrl" :size="240" color="#18a058" />
    </div>

    <div class="flex items-center justify-between text-center">
      <n-button
        strong
        secondary
        type="primary"
        size="large"
        @click="handleSecretUrlDetail"
        >查看秘钥</n-button
      >
      <n-button
        strong
        secondary
        type="error"
        size="large"
        @click="handleResetSecretUrl"
        >重置秘钥</n-button
      >
    </div>
  </div>
</template>
