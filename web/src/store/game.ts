import { defineStore } from "pinia";

export const useGameStore = defineStore("game", {
  persist: true,

  state: () => ({
    dimension: 4, //阶数
    colorPattern: 1, //颜色模式 0:层先 1:降阶
    gameMode: 0, //游戏模式
  }),

  getters: {
    getDimension: (state) => state.dimension,
    getGameMode: (state) => state.gameMode,
    getGameModeLabel: (state) => {
      if (state.gameMode === 0) return "练习";
      if (state.gameMode === 1) return "排位";
      if (state.gameMode === 2) return "对战";
    },
    getCellSet: (state) => {
      if (state.colorPattern === 0) {
        const redGroup = new Set<number>();
        const blueGroup = new Set<number>();
        const yellowGroup = new Set<number>();
        const indigoGroup = new Set<number>();
        const greenGroup = new Set<number>();
        const fuchsiaGroup = new Set<number>();

        for (let i = 1; i < state.dimension ** 2; i++) {
          if (Math.floor((i - 1) / state.dimension) + 1 === 1) {
            redGroup.add(i);
          }
          if (Math.floor((i - 1) / state.dimension) + 1 === 2) {
            blueGroup.add(i);
          }
          if (Math.floor((i - 1) / state.dimension) + 1 === 3) {
            yellowGroup.add(i);
          }
          if (Math.floor((i - 1) / state.dimension) + 1 === 4) {
            indigoGroup.add(i);
          }
          if (Math.floor((i - 1) / state.dimension) + 1 === 5) {
            greenGroup.add(i);
          }
          if (Math.floor((i - 1) / state.dimension) + 1 === 6) {
            fuchsiaGroup.add(i);
          }
        }

        return {
          redGroup,
          blueGroup,
          yellowGroup,
          indigoGroup,
          greenGroup,
          fuchsiaGroup,
        };
      }
      if (state.colorPattern === 1) {
        if (state.dimension === 3) {
          return {
            redGroup: new Set([1, 2, 3, 4, 7]),
            blueGroup: new Set([5, 6, 8]),
          };
        }
        if (state.dimension === 4) {
          return {
            redGroup: new Set([1, 2, 3, 4, 5, 9, 13]),
            blueGroup: new Set([6, 7, 8, 10, 14]),
            yellowGroup: new Set([11, 12, 15]),
          };
        }
        if (state.dimension === 5) {
          return {
            redGroup: new Set([1, 2, 3, 4, 5, 6, 11, 16, 21]),
            blueGroup: new Set([7, 8, 9, 10, 12, 17, 22]),
            indigoGroup: new Set([13, 14, 15, 18, 23]),
            yellowGroup: new Set([19, 20, 24]),
          };
        }
        if (state.dimension === 6) {
          return {
            redGroup: new Set([1, 2, 3, 4, 5, 6, 7, 13, 19, 25, 31]),
            blueGroup: new Set([8, 9, 10, 11, 12, 14, 20, 26, 32]),
            indigoGroup: new Set([15, 16, 17, 18, 21, 27, 33]),
            yellowGroup: new Set([22, 23, 24, 28, 34]),
            greenGroup: new Set([29, 30, 35]),
          };
        }
      }
    },
  },

  actions: {
    setDimension(dimension: number) {
      this.dimension = dimension;
    },
    setGameMode(gameMode: number) {
      this.gameMode = gameMode;
    },
    setColorPattern(colorPattern: number) {
      this.colorPattern = colorPattern;
    },
    getCellClass(cellValue: number) {
      if (this.getCellSet!.redGroup.has(cellValue)) return "bg-red-4 shadow";
      if (this.getCellSet!.blueGroup.has(cellValue)) return "bg-blue-4 shadow";
      if (this.getCellSet!.yellowGroup?.has(cellValue))
        return "bg-yellow-4 shadow";
      if (this.getCellSet!.indigoGroup?.has(cellValue))
        return "bg-indigo-4 shadow";
      if (this.getCellSet!.greenGroup?.has(cellValue))
        return "bg-green-4 shadow";
      if (this.getCellSet!.fuchsiaGroup?.has(cellValue))
        return "bg-orange-4 shadow";

      return "";
    },
  },
});
