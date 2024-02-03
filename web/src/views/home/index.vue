<script lang="ts" setup>
import { ref, computed } from "vue";
import TitleBar from "@/components/title-bar.vue";
import {
  getCellClass,
  getScramble,
  createHashMap,
  handleClick,
} from "@/common/game";
import { formatDurationInGame } from "@/utils/time";
import { useScrambleStore } from "@/store/scramble";
import { useGameStore } from "@/store/game";
import { useMessage } from "naive-ui";

const Message = useMessage();

const gameStore = useGameStore();
const scrambleStore = useScrambleStore();

const gameMap = ref([
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9, 10, 11, 12],
  [13, 14, 15, 0],
]);

const gameHashMap = ref(createHashMap(gameMap.value));

// 是否开始
const isStart = ref(false);
// 是否打乱
const isScramble = ref(false);
// 是否胜利
const isWin = ref(false);

// 游戏显示时间
const gameTime = computed((): string => {
  if (isStart.value) return formatDurationInGame(interval.value);
  if (isWin.value) return formatDurationInGame(endTime.value - startTime.value);
  return "00:00:000";
});
// 游戏步数
const gameStep = ref(0);
// 开始时间
const startTime = ref(0);
// 结束时间
const endTime = ref(0);
// 间隔时间
const interval = ref(0);
// 定时器
let timer: string | number | NodeJS.Timeout | undefined;
// 随机数
let idx = 0;

// 打乱顺序
let scrambleMap = [];
// 解法
let solution = [];

// 打乱
const handleScramble = () => {
  // if (scrambleStore.getScramble.scrambleStr !== "") {
  //   return message.error("请先完成本次还原");
  // }

  // 清除数据
  handleDataClear();

  // 打乱
  const { scramble, randomIdx } = getScramble(gameStore.getOrder, Date.now());

  // 设置打乱随机数
  idx = randomIdx;
  // 设置打乱顺序
  scrambleMap = scramble;
  // 持久化打乱信息
  scrambleStore.setScramble({
    scrambleMap: scramble,
    scrambleStr: scramble.join(","),
    scrambleIdx: idx,
  });

  // 将打乱顺序应用到游戏地图
  for (let i = 0; i < scramble.length; i++) {
    const row = Math.floor(i / 4);
    const column = i % 4;
    gameMap.value[row][column] = scramble[i];
    gameHashMap.value.set(scramble[i], { row, column });
  }

  // 设置打乱状态
  isScramble.value = true;
};

// 开始计时
const handleTimeStart = () => {
  timer = setInterval(() => {
    interval.value = Date.now() - startTime.value;
  }, 10);
};

// 处理点击
const handleTouch = (rowIndex: number, colIndex: number) => {
  if (isScramble.value) {
    isScramble.value = false;
    isStart.value = true;
    startTime.value = Date.now();
    handleTimeStart();
  }

  // 点击处理
  const { newSolution, newStepCount, map, hashMap } = handleClick(
    gameMap.value,
    gameHashMap.value,
    rowIndex,
    colIndex
  );

  // 更新数据
  gameMap.value = map;
  gameHashMap.value = hashMap;
  gameStep.value += newStepCount;
  newSolution.length && solution.push(...newSolution);

  // 判断是否胜利
  if (
    gameMap.value
      .flat()
      .slice(0, 15)
      .every((item, index) => item === index + 1) &&
    isStart.value
  ) {
    clearInterval(timer);
    isStart.value = false;
    isWin.value = true;
    endTime.value = Date.now();

    scrambleStore.clearScramble();
    scrambleStore.clearScrambleByShare();
  }
};

// 清除数据
const handleDataClear = () => {
  clearInterval(timer);
  isStart.value = false;
  isScramble.value = false;
  isWin.value = false;
  gameStep.value = 0;
  interval.value = 0;
  startTime.value = 0;
  endTime.value = 0;
  scrambleMap = [];
  solution = [];
};
</script>

<template>
  <div class="fixed inset-0 p-4 h-screen">
    <!-- 顶部 -->
    <div class="flex justify-between items-center mb-6">
      <title-bar title="练习" />
    </div>

    <!-- Map -->
    <div class="grid gap-2 p-1 rounded-md">
      <div
        class="grid grid-cols-4 gap-2"
        v-for="(row, rowIndex) in gameMap"
        :key="rowIndex"
      >
        <div
          class="aspect-square flex items-center justify-center rounded-md text-center text-5 font-bold"
          :class="item === 0 ? 'invisible' : getCellClass(item)"
          v-for="(item, colIndex) in row"
          :key="colIndex"
          @click="handleTouch(rowIndex, colIndex)"
        >
          {{ item }}
        </div>
      </div>
    </div>

    <!-- 时间与打乱 -->
    <div class="flex justify-around items-center w-full mt-10">
      <!-- 时间与步数 -->
      <div class="w-42 text-center tabular-nums">
        <div class="text-8 font-bold">{{ gameTime }}</div>
        <div class="text-4">步数: {{ gameStep }}</div>
      </div>

      <n-button
        class="w-30 h-15 text-5 shadow"
        strong
        secondary
        type="primary"
        @click="handleScramble"
      >
        打乱
      </n-button>
    </div>
  </div>
</template>
