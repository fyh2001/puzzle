<template>
  <div>
    <n-config-provider :theme="theme" :theme-overrides="themeColor">
      <n-global-style />
      <n-message-provider container-class="message" :max="2">
        <n-dialog-provider>
          <slot />
        </n-dialog-provider>
      </n-message-provider>
    </n-config-provider>
  </div>
</template>

<script lang="ts" setup>
import { computed } from "vue";
import { darkTheme, lightTheme } from "naive-ui";
import { useThemeStore } from "@/store/theme";
import { themeProvider } from "@/theme";
const themeStore = useThemeStore();

themeStore.fetchLoadedThemes();

const theme = computed(() => (themeStore.darkMode ? darkTheme : lightTheme));
const themeColor = computed(() => {
  /** 默认主题（这里采用空对象，也手动导入指定的主题 */
  const defaultTheme = {
    light: {},
    dark: {},
  };
  const resultTheme =
    themeStore.loadedThemesMap.get(themeStore.themeColor) || defaultTheme;
  return themeProvider(resultTheme[themeStore.darkMode ? "dark" : "light"]);
});
</script>

<style></style>
