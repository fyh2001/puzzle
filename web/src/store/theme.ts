import { defineStore } from "pinia";
import { Theme } from "@/types/common";

export const useThemeStore = defineStore("theme", {
  persist: true,

  state: () => ({
    // 深色模式/浅色模式
    darkMode: false,
    // 主题色
    themeColor: "green",

    // 已加载的主题
    loadedThemes: [] as Theme[],
  }),

  getters: {
    // 深色模式/浅色模式
    getDarkMode: (state) => state.darkMode,
    // 主题色
    getThemeColor: (state) => state.themeColor,
    // 已加载的主题的name映射
    loadedThemesMap: (state) =>
      new Map(state.loadedThemes.map((item) => [item.name, item])),
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

    // 同步已加载的主题
    async fetchLoadedThemes() {
      // 动态加载 theme 文件夹下的全部主题，并转换为映射
      const themeModules = import.meta.glob("@/theme/lib/*.ts");
      Promise.all(Object.values(themeModules).map((val) => val())).then(
        (themes) => {
          this.loadedThemes = themes.map((item: any) => item.default);
        }
      );
    },
  },
});
