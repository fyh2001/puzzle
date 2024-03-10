<script lang="ts" setup>
import { ChevronRightRound } from "@vicons/material";
import { defaultAvatar } from "@/config";
import { useDialog, useMessage } from "naive-ui";
import router from "@/routers";
import { useEventListener } from "@vueuse/core";
import type { UserModel } from "@/types/user";

const props = defineProps<{
  label: string;
  content?: string;
  avatar?: string;
}>();

const emit = defineEmits<{
  (e: "update", form: UserModel): void;
}>();

const slots = defineSlots<{
  dialog(props: { update: (form: UserModel) => void }): any;
}>();

const dialog = useDialog();
const message = useMessage();

// 公共的更新函数
const update = (form: Record<string, any>) => emit("update", form);

const showDialog = () => {
  if (!slots.dialog) {
    return message.warning("暂未开放");
  }

  handleBackEventLister();

  dialog.warning({
    title: props.label,
    style: { borderRadius: "0.75rem" },
    transformOrigin: "center",
    maskClosable: false,
    content: () => slots.dialog({ update }),
    onClose: () => {
      window.history.back();
    },
  });
};

// 监听返回事件
const handleBackEventLister = () => {
  // 防止刷新页面后，路由不变
  window.history.pushState(null, "", `#${router.currentRoute.value.path}`);
  useEventListener(
    "popstate",
    () => {
      dialog.destroyAll();
    },
    { once: true }
  );
};
</script>

<template>
  <div class="flex justify-between items-center" @click="showDialog">
    <div>{{ label }}</div>
    <div class="flex justify-between items-center gap-3">
      <n-avatar
        class="shadow-lg"
        round
        object-fit="cover"
        :size="52"
        :src="avatar"
        :fallback-src="defaultAvatar"
        v-if="avatar"
      />
      <div class="text-gray" v-if="content">{{ content }}</div>
      <n-icon class="text-gray" :size="26" :component="ChevronRightRound" />
    </div>
  </div>
</template>
