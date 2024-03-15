<script lang="ts" setup>
import { computed } from "vue";
import { useCommonStore } from "@/store/common";

const commonStore = useCommonStore();

interface TextSize {
  size: string;
  language: string;
}

const props = withDefaults(
  defineProps<{
    title: string;
    textSize?: TextSize[];
    textColor?: string;
  }>(),
  {
    title: "",
    textColor: "var(--neutralTextBase)",
  }
);

const textSizeComputed = computed(() => {
  return (
    props.textSize?.find((item) => item.language === commonStore.getLanguage)
      ?.size || "26px"
  );
});
</script>

<template>
  <div>
    <n-el :style="{ fontSize: textSizeComputed, color: textColor }">{{
      props.title
    }}</n-el>
  </div>
</template>
