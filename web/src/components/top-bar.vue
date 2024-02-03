<script lang="ts" setup>
import ChevronLeftRound from "@vicons/material/ChevronLeftRound";
import router from "@/routers/index";

defineProps<{
  title: string;
}>();
</script>

<template>
  <div>
    <div class="flex items-center justify-between">
      <div class="flex items-center" @click="router.back()">
        <n-icon class="text-8" :component="ChevronLeftRound" />
        <div class="text-4.2" v-if="title">{{ title }}</div>
      </div>
      <slot name="right" />
    </div>
  </div>
</template>
