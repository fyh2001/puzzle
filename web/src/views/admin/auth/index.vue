<script lang="ts" setup>
import { ref } from "vue";
import { useAdminStore } from "@/store/admin";
import { useMessage } from "naive-ui";
import router from "@/routers";

const Message = useMessage();
const adminStore = useAdminStore();

const code = ref("");

const onFinish = (code: string) => {
  adminStore
    .auth({ code })
    .then(() => {
      Message.success("登录成功");
      setTimeout(() => {
        router.push({ name: "Admin" });
      }, 1000);
    })
    .catch(() => {
      Message.error("登录失败");
    });
};
</script>

<template>
  <div class="w-screen h-screen bg-white">
    <a-verification-code
      class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2"
      v-model="code"
      style="width: 300px"
      @finish="onFinish"
    />
  </div>
</template>

<style scoped>
:deep body {
  background-color: white;
}
</style>
