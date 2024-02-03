import { defineStore } from "pinia";

export const useThemeStore = defineStore("theme", {
  persist: true,

  state: () => ({
    // 深色模式/浅色模式
    darkMode: false,
    // 主题色
    themeColor: "green",
  }),

  getters: {
    // 深色模式/浅色模式
    getDarkMode: (state) => state.darkMode,
    // 主题色
    getThemeColor: (state) => state.themeColor,
  },

  actions: {
    // 打开深色模式
    openDarkMode() {
      this.darkMode = true;
    },

    // 关闭深色模式
    closeDarkMode() {
      this.darkMode = false;
    },

    // 切换主题色
    changeThemeColor(color: string) {
      this.themeColor = color;
    },
  },
});
