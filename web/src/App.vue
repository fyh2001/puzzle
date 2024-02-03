<script lang="ts" setup>
import GlobalProvider from "@/components/global-provider.vue";
import { useThemeStore } from "@/store/theme";

import { onMounted } from "vue";

const themeStore = useThemeStore();

onMounted(() => {});
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
