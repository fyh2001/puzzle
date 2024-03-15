<script lang="ts" setup>
import GlobalProvider from "@/components/global-provider.vue";
import GlobalInit from "@/components/global-init.vue";
import { useThemeStore } from "@/store/theme";

const themeStore = useThemeStore();
</script>

<template>
  <global-provider>
    <global-init />
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
