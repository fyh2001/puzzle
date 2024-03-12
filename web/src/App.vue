<script lang="ts" setup>
import GlobalProvider from "@/components/global-provider.vue";
import polling from "@/components/polling.vue";
import { onMounted } from "vue";
import { useThemeStore } from "@/store/theme";

const themeStore = useThemeStore();

onMounted(() => {});
</script>

<template>
  <global-provider>
    <polling />
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
