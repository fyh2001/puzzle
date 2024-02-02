import { defineStore } from "pinia";

export const useGameStore = defineStore("game", {
  persist: true,

  state: () => ({
    // 阶数
    order: 4,
  }),

  getters: {
    // 获取阶数
    getOrder: (state) => state.order,
  },

  actions: {
    // 设置阶数
    setOrder(order: number) {
      this.order = order;
    },
  },
});
