import { defineStore } from "pinia";
import type { ScrambleModel } from "@/store/scramble/types";

export const useScrambleStore = defineStore("scramble", {
  persist: true,

  state: () => ({
    scramble: <ScrambleModel>{},
    scrambleByShare: <ScrambleModel>{},
  }),

  getters: {
    getScramble: (state) => state.scramble,
    getScrambleByShare: (state) => state.scrambleByShare,
  },

  actions: {
    setScramble(scramble: ScrambleModel) {
      this.scramble = scramble;
    },

    setScrambleByShare(scramble: ScrambleModel) {
      this.scrambleByShare = scramble;
    },

    clearScramble() {
      this.scramble = <ScrambleModel>{};
    },

    clearScrambleByShare() {
      this.scrambleByShare = <ScrambleModel>{};
    },
  },
});
