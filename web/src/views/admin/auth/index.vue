<script lang="ts" setup>
import { ref } from "vue";
import { useAdminStore } from "@/store/admin";
import { useMessage } from "naive-ui";
import router from "@/routers";

const Message = useMessage();
const adminStore = useAdminStore();

const inputValue = ref("");
const isCodeError = ref(false);

const onFinish = (code: string) => {
  adminStore.auth({ code }).then((res) => {
    if (res) {
      Message.success("登录成功");
      return setTimeout(() => {
        router.push({ name: "Admin" });
      }, 1000);
    }

    inputValue.value = "";
    isCodeError.value = true;
    Message.error("登录失败");
  });
};
</script>

<template>
  <div class="w-screen h-screen bg-white">
    <a-verification-code
      class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2"
      v-model="inputValue"
      :error="isCodeError"
      style="width: 300px"
      @change="
        () => {
          isCodeError && (isCodeError = false);
        }
      "
      @finish="onFinish"
    />
  </div>
</template>

<style scoped>
:deep body {
  background-color: white;
}
</style>
