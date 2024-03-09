<script lang="ts" setup>
import { ChevronRightRound } from "@vicons/material";
import { defaultAvatar } from "@/config";
import { useDialog, useMessage } from "naive-ui";

const props = defineProps<{
  label: string;
  content?: string;
  avatar?: string;
}>();

const emit = defineEmits<{
  (e: "update", form: Record<string, any>): void;
}>();

const slots = defineSlots<{
  dialog(props: { update: (form: Record<string, any>) => void }): any;
}>();

const dialog = useDialog();
const message = useMessage();

// 公共的更新函数
const update = (form: Record<string, any>) => emit("update", form);

const showDialog = () => {
  if (!slots.dialog) {
    message.warning("暂未开放");
    return;
  }

  dialog.warning({
    title: props.label,
    content: () => slots.dialog({ update }),
  });
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
