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
  ColorLensRound,
  GridOnRound,
  AutoAwesomeMotionRound,
} from "@vicons/material";
import { ref, computed, watch } from "vue";
import TitleBar from "@/components/title-bar.vue";
import Dropdown from "@/components/dropdown.vue";
import {
  shuffle,
  createHashMap,
  handleClick,
  transformArray,
} from "@/common/game";
import { formatDurationInGame } from "@/utils/time";
import { useScrambleStore } from "@/store/scramble";
import { useGameStore } from "@/store/game";
import { useMessage } from "naive-ui";
import { recordRequest } from "@/api/methods/record";
import { scrambleRequest } from "api/scramble";

const Message = useMessage();

const gameStore = useGameStore();
const scrambleStore = useScrambleStore();

// 下拉选项
const dropdownOptions = [
  {
    label: "游戏模式",
    key: "mode",
    disabled: false,
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
            gameStore.setGameMode(1);
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
            gameStore.setGameMode(2);
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
          },
        },
      },
    ],
  },
  {
    label: "颜色模式",
    key: "colorPattern",
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={ColorLensRound} />
        </n-el>
      );
    },
    children: [
      {
        label: "层先",
        key: "layerFirst",
        disabled: false,
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon size="18" component={GridOnRound} />
            </n-el>
          );
        },
        props: {
          onClick: () => {
            gameStore.setColorPattern(0);
          },
        },
      },
      {
        label: "降阶",
        key: "decreaseDimension",
        disabled: false,
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon size="18" component={AutoAwesomeMotionRound} />
            </n-el>
          );
        },
        props: {
          onClick: () => {
            gameStore.setColorPattern(1);
          },
        },
      },
    ],
  },
];

const gameMap = ref([
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9, 10, 11, 12],
  [13, 14, 15, 0],
]);
const gameHashMap = ref(createHashMap(gameMap.value));

// 是否加载中
const isLoding = ref(false);
// 打乱按钮是否禁用
const btnDisabled = computed(() => {
  return gameStore.getGameMode === 2 && isScramble.value;
});

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
let scrambleMap: number[] | undefined = [];
// 解法
let solution: string[] = [];

// 应用打乱顺序到地图
const handleScrambleToMap = (scramble: number[], dimension: number) => {
  for (let i = 0; i < scramble.length; i++) {
    const row = Math.floor(i / dimension);
    const column = i % dimension;
    gameMap.value[row][column] = scramble[i];
    gameHashMap.value.set(scramble[i], { row, column });
  }
};

// 设置打乱状态
const setScrambleStatus = (val: boolean) => {
  isScramble.value = val;
};

// 打乱
const handleScramble = async () => {
  // 清除数据
  gameStore.getGameMode === 1 && handleDataClear();

  // 练习模式
  if (gameStore.getGameMode === 1) {
    idx = Date.now(); // 设置打乱随机数
    scrambleMap = shuffle(gameStore.getDimension, idx); // 打乱顺序
  }

  // 排位模式
  if (gameStore.getGameMode === 2) {
    let { scrambleArr, scrambleIdx } = await getRankScramble();

    if (scrambleArr!.length === 0) return;

    handleDataClear();

    scrambleMap = scrambleArr;
    idx = scrambleIdx;
  }

  handleScrambleToMap(scrambleMap!, gameStore.getDimension);

  setScrambleStatus(true);
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
  const { newSolution, newStepCount, newMap, hashMap } = handleClick(
    gameMap.value,
    gameHashMap.value,
    rowIndex,
    colIndex
  );

  // 更新数据
  gameMap.value = newMap;
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
  const {
    data: { code, data: resp, msg },
  } = await recordRequest.insert({
    dimension: gameStore.getDimension,
    type: gameStore.getGameMode,
    duration: endTime.value - startTime.value,
    step: gameStep.value,
    scramble: scrambleMap!.join(","),
    solution: solution.join(","),
    idx: idx,
  });

  if (code === 200) {
    return Message.success(resp);
  }

  Message.error(msg);
};

// 初始化练习地图
const initMap = () => {
  const newGameMap = Array.from(
    { length: gameStore.getDimension * gameStore.getDimension },
    (_, index) => {
      return index + 1;
    }
  );

  newGameMap[newGameMap.length - 1] = 0;

  gameMap.value = transformArray(newGameMap, gameStore.getDimension);
  gameHashMap.value = createHashMap(gameMap.value);
  handleDataClear();
};

// 初始化排位地图
const initRankMap = async () => {
  const {
    data: { code, msg, data: scrambleData },
  } = await scrambleRequest.getUserScramble({
    dimension: gameStore.getDimension,
  });

  if (code === 200) {
    if (scrambleData.idx === 0) return initMap();

    const scramble = shuffle(scrambleData.dimension, scrambleData.idx); // 打乱顺序

    // 设置打乱顺序
    scrambleMap = scramble;
    // 设置idx
    idx = scrambleData.idx;

    // 将打乱顺序应用到游戏地图
    handleScrambleToMap(scramble, scrambleData.dimension);

    setScrambleStatus(true);
  } else {
    Message.error(msg);
  }
};

// 获取排位打乱
const getRankScramble = async () => {
  isLoding.value = true;
  const {
    data: { code, msg, data: scrambleData },
  } = await scrambleRequest.getNewScramble({
    dimension: gameStore.getDimension,
  });

  isLoding.value = false;
  if (code === 200) {
    const scrambleArr = shuffle(scrambleData.dimension, scrambleData.idx);
    return { scrambleArr, scrambleIdx: scrambleData.idx };
  } else {
    Message.error(msg);
    return { scrambleArr: [], scrambleIdx: 0 };
  }
};

// 监听游戏维度变化
watch(
  () => [gameStore.getDimension, gameStore.gameMode],
  () => {
    initMap();
    if (gameStore.gameMode === 2) {
      initRankMap();
    }
  },
  { immediate: true }
);
</script>

<template>
  <div class="fixed inset-0 p-4 h-screen">
    <!-- 顶部 -->
    <div class="flex justify-between items-center mb-6">
      <title-bar :title="gameStore.getGameModeLabel" />
      <dropdown content="功能" :options="dropdownOptions" :showDivider="true" />
    </div>

    <n-spin :show="isLoding" description="正在获取打乱" size="large">
      <!-- Map -->
      <div class="grid gap-2 p-1 rounded-md">
        <div
          class="grid gap-2"
          :style="`grid-template-columns: repeat(${gameStore.getDimension}, 1fr);`"
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
          :disabled="btnDisabled"
          @click="handleScramble"
        >
          打乱
        </n-button>
      </div>
    </n-spin>
  </div>
</template>
