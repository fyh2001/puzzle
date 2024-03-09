<script lang="ts" setup>
import { ChevronRightRound } from "@vicons/material";
import { defaultAvatar } from "@/config";
defineProps<{
  label: string;
  content?: string;
  avatar?: string;
}>();
</script>

<template>
  <div class="flex justify-between items-center">
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

<style scoped></style>
