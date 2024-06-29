import { defineStore } from "pinia";

export const useCommonStore = defineStore("common", {
  persist: true,

  state: () => ({
    language: "cn",
  }),

  getters: {
    getLanguage: (state) => state.language,
  },

  actions: {
    changeLanguage(language: string) {
      this.language = language;
    },
  },
});
