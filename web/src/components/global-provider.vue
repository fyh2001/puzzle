<template>
  <div>
    <n-config-provider :theme="theme" :theme-overrides="themeColor">
      <n-global-style />
      <n-message-provider :max="2">
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
import { greenLightTheme, greenDarkTheme } from "@/theme/green";
import { indigoLightTheme, indigoDarkTheme } from "@/theme/indigo";
import { roseLightTheme, roseDarkTheme } from "@/theme/rose";
import { amberLightTheme, amberDarkTheme } from "@/theme/amber";
import { skyLightTheme, skyDarkTheme } from "@/theme/sky";
import { useThemeStore } from "@/store/theme";

const themeStore = useThemeStore();

const theme = computed(() => {
  return themeStore.darkMode ? darkTheme : lightTheme;
});
const themeColor = computed(() => {
  switch (themeStore.themeColor) {
    case "green":
      return themeStore.darkMode ? greenDarkTheme : greenLightTheme;
    case "indigo":
      return themeStore.darkMode ? indigoDarkTheme : indigoLightTheme;
    case "rose":
      return themeStore.darkMode ? roseDarkTheme : roseLightTheme;
    case "amber":
      return themeStore.darkMode ? amberDarkTheme : amberLightTheme;
    case "sky":
      return themeStore.darkMode ? skyDarkTheme : skyLightTheme;
    default:
      break;
  }
});
</script>
