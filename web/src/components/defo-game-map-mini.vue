<script setup lang="ts">
import {
  onMounted,
  onBeforeUnmount,
  ref,
  defineComponent,
  computed,
  watch,
} from "vue";
import { SwipeLeftRound } from "@vicons/material";
import { useGameStore } from "@/store/game";
import {
  getGameMapByScramble,
  handleClick,
  createHashMap,
} from "@/common/game";
import { useI18n } from "vue-i18n";

defineComponent({
  name: "DefoGameMapMini",
});

const { t } = useI18n();

const gameStore = useGameStore();

const props = withDefaults(
  defineProps<{
    scramble: string;
    dimension: any;
    solution?: string;
    buttonAlign?: string;
    buttonSize?: string;
    buttonClass?: string;
  }>(),
  {
    buttonAlign: "center",
    buttonSize: "large",
    buttonClass: "",
  }
);

// 是否正在加载
const isLoding = ref(false);

const gameMap = ref<number[][]>([]);
const gameMapList = ref<Array<number[][]>>([]);
const gameMapListIndex = ref(0); // 当前步数

// 通过解法获取每次的Map
const getSolutionCase = () => {
  if (!props.solution) return;

  isLoding.value = true;

  let currentMap = getGameMapByScramble(props.scramble, props.dimension);
  let currentHashMap = createHashMap(currentMap);
  const solution = props.solution!.split(",");

  gameMapList.value.push(currentMap);

  solution.forEach((item) => {
    const { newMap, hashMap } = handleClick(
      currentMap,
      currentHashMap,
      undefined,
      undefined,
      parseInt(item)
    );

    currentMap = newMap;
    currentHashMap = hashMap;

    gameMapList.value.push(currentMap);
  });

  isLoding.value = false;
};

// 下一步按钮label
const nextLabel = computed(() => {
  return gameMapListIndex.value === gameMapList.value.length - 1
    ? t("miniGameMap.reset")
    : t("miniGameMap.next");
});
// 当前自动播放按钮label
const autoPlayLabel = computed(() => {
  return timer.value ? t("miniGameMap.stopPlay") : t("miniGameMap.autoPlay");
});
// 当前自动播放的定时器
let timer: any = ref(null);
// 当前自动播放速度
const speed = ref(500);
// 自动播放速度label
const speedLabelOptions: any = {
  1000: "0.5x",
  500: "1x",
  250: "2x",
  125: "4x",
};
// 自动播放
const handleAutoPlay = () => {
  if (timer.value) {
    clearInterval(timer.value);
    timer.value = null;
  } else {
    timer.value = setInterval(() => {
      gameMapListIndex.value = gameMapListIndex.value + 1;
      if (gameMapListIndex.value === gameMapList.value.length - 1) {
        clearInterval(timer.value);
        timer.value = null;
      }
    }, speed.value);
  }
};
// 清除自动播放
const handleAutoPlayClear = () => {
  if (timer.value) {
    clearInterval(timer.value);
    timer.value = null;
  }
};

onMounted(() => {
  gameMap.value = getGameMapByScramble(props.scramble, props.dimension);
  getSolutionCase();
});

onBeforeUnmount(() => {
  if (timer.value) {
    clearInterval(timer.value);
    timer.value = null;
  }
});

watch(
  () => speed.value as number,
  () => {
    if (timer.value) {
      clearInterval(timer.value);
      timer.value = setInterval(() => {
        gameMapListIndex.value = gameMapListIndex.value + 1;
        if (gameMapListIndex.value === gameMapList.value.length - 1) {
          clearInterval(timer.value);
          timer.value = null;
        }
      }, speed.value);
    }
  }
);
</script>

<template>
  <div w-full h-full>
    <n-tag
      :bordered="false"
      type="primary"
      class="mb-4 rounded text-4 font-bold font-mono"
      v-if="solution"
    >
      <template #icon>
        <n-icon :component="SwipeLeftRound" />
      </template>
      {{ t("miniGameMap.step") }}: {{ gameMapListIndex }} /
      {{ gameMapList.length - 1 }}
    </n-tag>
    <!-- Map -->
    <div class="grid gap-2 p-1 w-full h-full rounded-md" v-if="!solution">
      <div
        class="grid gap-2"
        :style="`grid-template-columns: repeat(${dimension}, 1fr);`"
        v-for="(row, rowIndex) in gameMap"
        :key="rowIndex"
      >
        <div
          class="aspect-square flex items-center justify-center rounded-md text-center text-5 font-bold"
          :class="
            item === 0 ? 'invisible' : gameStore.getCellClass(item, dimension)
          "
          v-for="(item, colIndex) in row"
          :key="colIndex"
        >
          {{ item }}
        </div>
      </div>
    </div>
    <!-- Map 与 解法 -->
    <div class="grid gap-2 p-1 w-full h-full rounded-md" v-else>
      <div
        class="grid gap-2"
        :style="`grid-template-columns: repeat(${dimension}, 1fr);`"
        v-for="(row, rowIndex) in gameMapList[gameMapListIndex]"
        :key="rowIndex"
      >
        <div
          class="aspect-square flex items-center justify-center rounded-md text-center text-5 font-bold"
          :class="
            item === 0 ? 'invisible' : gameStore.getCellClass(item, dimension)
          "
          v-for="(item, colIndex) in row"
          :key="colIndex"
        >
          {{ item }}
        </div>
      </div>
    </div>
    <!-- 按钮 -->
    <div v-if="solution">
      <div class="flex justify-between items-center mt-4">
        <n-button
          :class="buttonClass"
          strong
          secondary
          type="primary"
          :size="buttonSize"
          @click="handleAutoPlay"
          :disabled="gameMapListIndex === gameMapList.length - 1"
          >{{ autoPlayLabel }}</n-button
        >
        <n-dropdown
          size="large"
          placement="top-start"
          :class="{ invisible: !timer }"
          :options="[
            { label: '0.5x', key: 1000 },
            { label: '1x', key: 500 },
            { label: '2x', key: 250 },
            { label: '4x', key: 125 },
          ]"
          @select="($event: number) => (speed = $event)"
        >
          <n-button
            strong
            secondary
            :size="buttonSize"
            type="primary"
            :class="[buttonClass, { invisible: !timer }]"
            >{{ t("miniGameMap.rate") }}
            {{ speedLabelOptions[speed] }}</n-button
          >
        </n-dropdown>
        <n-button
          :class="buttonClass"
          strong
          secondary
          type="primary"
          :size="buttonSize"
          @click="
            () => {
              handleAutoPlayClear();
              gameMapListIndex = gameMapListIndex - 1;
            }
          "
          :disabled="gameMapListIndex === 0"
          >{{ t("miniGameMap.last") }}</n-button
        >
        <n-button
          :class="buttonClass"
          strong
          secondary
          type="primary"
          :size="buttonSize"
          @click="
            () => {
              handleAutoPlayClear();
              gameMapListIndex = gameMapListIndex + 1;
              if (gameMapListIndex === gameMapList.length) gameMapListIndex = 0;
            }
          "
          >{{ nextLabel }}</n-button
        >
      </div>
    </div>
  </div>
</template>
