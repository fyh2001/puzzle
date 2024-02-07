<script lang="tsx" setup>
import {
  GamesRound,
  MovingRound,
  WorkspacePremiumRound,
  CableRound,
  Looks3Outlined,
  Looks4Outlined,
  Looks5Outlined,
  Looks6Outlined,
  DashboardCustomizeRound,
} from "@vicons/material";
import { ref, computed, watch } from "vue";
import TitleBar from "@/components/title-bar.vue";
import Dropdown from "@/components/dropdown.vue";
import {
  getScramble,
  createHashMap,
  handleClick,
  transformArray,
} from "@/common/game";
import { formatDurationInGame } from "@/utils/time";
import { useScrambleStore } from "@/store/scramble";
import { useGameStore } from "@/store/game";
import { useMessage } from "naive-ui";
import { recordRequest } from "@/api/methods/record";
const Message = useMessage();

const gameStore = useGameStore();
const scrambleStore = useScrambleStore();

// 下拉选项
const dropdownOptions = [
  {
    label: "游戏模式",
    key: "mode",
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={GamesRound} />
        </n-el>
      );
    },
    children: [
      {
        label: "练习",
        key: "practice",
        disabled: false,
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon size="18" component={MovingRound} />
            </n-el>
          );
        },
        props: {
          onClick: () => {
            gameMode.value = "practice";
          },
        },
      },
      {
        label: "排位",
        key: "ranked",
        disabled: false,
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon size="18" component={WorkspacePremiumRound} />
            </n-el>
          );
        },
        props: {
          onClick: () => {
            gameMode.value = "ranked";
          },
        },
      },
      {
        label: "对战",
        key: "battle",
        disabled: true,
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon size="18" component={CableRound} />
            </n-el>
          );
        },
        // props: {
        //   onClick: () => {
        //     gameMode.value = "battle";
        //   },
        // },
      },
    ],
  },
  {
    label: "阶数选择",
    key: "dimension",
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={DashboardCustomizeRound} />
        </n-el>
      );
    },
    children: [
      {
        label: "3x3",
        key: "3",
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon size="18" component={Looks3Outlined} />
            </n-el>
          );
        },
        props: {
          onClick: () => {
            gameStore.setDimension(3);
            gameDimension.value = 3;
          },
        },
      },
      {
        label: "4x4",
        key: "4",
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon size="18" component={Looks4Outlined} />
            </n-el>
          );
        },
        props: {
          onClick: () => {
            gameStore.setDimension(4);
            gameDimension.value = 4;
          },
        },
      },
      {
        label: "5x5",
        key: "5",
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon size="18" component={Looks5Outlined} />
            </n-el>
          );
        },
        props: {
          onClick: () => {
            gameStore.setDimension(5);
            gameDimension.value = 5;
          },
        },
      },
      {
        label: "6x6",
        key: "6",
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon size="18" component={Looks6Outlined} />
            </n-el>
          );
        },
        props: {
          onClick: () => {
            gameStore.setDimension(6);
            gameDimension.value = 6;
          },
        },
      },
    ],
  },
];

// 当前游戏模式
const gameMode = ref("practice");
// 当前游戏维度
const gameDimension = ref(gameStore.dimension);

// 当前游戏模式
const currentGameModeValue = computed(() => {
  return {
    practice: 0,
    ranked: 1,
    battle: 2,
  }[gameMode.value];
});
// 当前游戏模式标签
const currentGameModeLabel = computed(
  () =>
    ({
      practice: "练习",
      ranked: "排位",
      battle: "对战",
    }[gameMode.value])
);

const gameMap = ref([
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9, 10, 11, 12],
  [13, 14, 15, 0],
]);

const gameHashMap = ref(createHashMap(gameMap.value));

// 监听游戏维度变化
watch(
  () => gameDimension.value,
  (newVal) => {
    const newGameMap = Array.from({ length: newVal * newVal }, (_, index) => {
      return index + 1;
    });

    newGameMap[newGameMap.length - 1] = 0;

    gameMap.value = transformArray(newGameMap, newVal);
    gameHashMap.value = createHashMap(gameMap.value);
  },
  { immediate: true }
);

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
let scrambleMap: number[] = [];
// 解法
let solution: string[] = [];

function shuffle(array: number[], seed: number) {
  let currentIndex = array.length,
    temporaryValue,
    randomIndex;

  // 设置随机数种子
  const seededRandom = (min: number, max: number) => {
    let x = Math.sin(seed++) * 10000;
    return Math.floor((x - Math.floor(x)) * (max - min + 1)) + min;
  };

  // 当还有未洗牌的元素时
  while (currentIndex !== 0) {
    // 随机选取一个未洗牌的元素
    randomIndex = seededRandom(0, currentIndex - 1);
    currentIndex--;

    // 交换它与当前元素
    temporaryValue = array[currentIndex];
    array[currentIndex] = array[randomIndex];
    array[randomIndex] = temporaryValue;
  }

  return array;
}

// 打乱
const handleScramble = () => {
  // if (scrambleStore.getScramble.scrambleStr !== "") {
  //   return message.error("请先完成本次还原");
  // }

  // 清除数据
  handleDataClear();

  // 打乱
  const scramble = shuffle(gameMap.value.flat(), 100);

  // 将打乱顺序应用到游戏地图
  for (let i = 0; i < scramble.length; i++) {
    const row = Math.floor(i / gameDimension.value);
    const column = i % gameDimension.value;
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
  newSolution.length && solution.push(newSolution);

  // 判断是否胜利
  if (
    gameMap.value
      .flat()
      .slice(0, gameMap.value.flat().length - 1)
      .every((item, index) => item === index + 1) &&
    isStart.value
  ) {
    clearInterval(timer);
    isStart.value = false;
    isWin.value = true;
    endTime.value = Date.now();

    scrambleStore.clearScramble();
    scrambleStore.clearScrambleByShare();

    handleRecordUpload();
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

// 上传记录
const handleRecordUpload = async () => {
  console.log({
    dimension: gameDimension.value,
    type: currentGameModeValue.value,
    duration: endTime.value - startTime.value,
    step: gameStep.value,
    scramble: scrambleMap.join(","),
    solution: solution.join(","),
    idx: idx,
  });

  console.log(solution);

  const {
    data: { code, data: resp },
  } = await recordRequest.insert({
    dimension: gameDimension.value,
    type: currentGameModeValue.value,
    duration: endTime.value - startTime.value,
    step: gameStep.value,
    scramble: scrambleMap.join(","),
    solution: solution.join(","),
    idx: idx,
  });

  if (code === 200) {
    return Message.success(resp);
  }

  Message.error(resp);
};
</script>

<template>
  <div class="fixed inset-0 p-4 h-screen">
    <!-- 顶部 -->
    <div class="flex justify-between items-center mb-6">
      <title-bar :title="currentGameModeLabel" />
      <dropdown content="功能" :options="dropdownOptions" :showDivider="true" />
    </div>

    <!-- Map -->
    <div class="grid gap-2 p-1 rounded-md">
      <div
        class="grid gap-2"
        :style="`grid-template-columns: repeat(${gameDimension}, 1fr);`"
        v-for="(row, rowIndex) in gameMap"
        :key="rowIndex"
      >
        <div
          class="aspect-square flex items-center justify-center rounded-md text-center text-5 font-bold"
          :class="item === 0 ? 'invisible' : gameStore.getCellClass(item)"
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
