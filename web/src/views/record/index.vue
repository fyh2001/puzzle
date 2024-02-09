<script lang="tsx" setup>
import { ref, computed, markRaw } from "vue";
import {
  WorkspacePremiumRound,
  Looks3Outlined,
  Looks4Outlined,
  Looks5Outlined,
  Looks6Outlined,
  DashboardCustomizeRound,
  RocketLaunchOutlined,
  Filter5Round,
  TimerOutlined,
  SwipeLeftRound,
  Filter9PlusRound,
  CalendarMonthRound,
} from "@vicons/material";
import TitleBar from "@/components/title-bar.vue";
import Dropdown from "@/components/dropdown.vue";
import recordPerson from "@/views/record/components/record-person.vue";

import { useGameStore } from "@/store/game";

const tableOptions = {
  mine: {
    title: "我的",
    view: recordPerson,
  },
  all: {
    title: "总排名",
    children: {
      bestSingle: {
        title: "最佳单次",
        icon: markRaw(TimerOutlined),
        view: "bbb",
      },
      bestAverage5: {
        title: "最佳5次平均",
        icon: markRaw(Filter5Round),
      },
      bestAverage12: {
        title: "最佳12次平均",
        icon: markRaw(Filter9PlusRound),
      },
      bestStepCount: {
        title: "最佳步数",
        icon: markRaw(SwipeLeftRound),
      },
    },
  },
  weekly: {
    title: "周排名",
    children: {
      bestSingle: {
        title: "最佳单次",
        icon: markRaw(TimerOutlined),
      },
      bestAverage5: {
        title: "最佳5次平均",
        icon: markRaw(Filter5Round),
      },
      bestAverage12: {
        title: "最佳12次平均",
        icon: markRaw(Filter9PlusRound),
      },
      bestStepCount: {
        title: "最佳步数",
        icon: markRaw(SwipeLeftRound),
      },
    },
  },
};
const currentTableOptions = ref("mine");
const currentChildOption = ref("bestSingle");

const gameStore = useGameStore();
// 功能下拉框选项
const options = [
  // 阶数
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
  // 我的记录
  {
    label: "我的记录",
    key: "record",
    props: {
      onClick: () => {
        currentTableOptions.value = "mine";
      },
    },
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={RocketLaunchOutlined} />
        </n-el>
      );
    },
  },
  // 总排名
  {
    label: "总排名",
    key: "rank",
    disabled: false,
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={WorkspacePremiumRound} />
        </n-el>
      );
    },
    children: Object.entries(tableOptions.all.children).map(([key, val]) => ({
      label: val.title,
      key: key,
      props: {
        onClick: () => {
          currentTableOptions.value = "all";
          currentChildOption.value = key;
        },
      },
      icon: () => (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={val.icon} />
        </n-el>
      ),
    })),
  },
  // 周排名
  {
    label: "周排名",
    key: "weekRank",
    disabled: true,
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={CalendarMonthRound} />
        </n-el>
      );
    },
    children: [
      {
        label: "最佳单次",
        key: "weekBestSingle",
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon size="18" component={TimerOutlined} />
            </n-el>
          );
        },
      },
      {
        label: "最佳5次平均",
        key: "weekBestAverage5",
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon size="18" component={Filter5Round} />
            </n-el>
          );
        },
      },
      {
        label: "最佳12次平均",
        key: "weekBestAverage12",
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon size="18" component={Filter9PlusRound} />
            </n-el>
          );
        },
      },
      {
        label: "最佳步数",
        key: "weekBestStepCount",
        icon: () => {
          return (
            <n-el class="flex items-center" style="color: var(--primary-color)">
              <n-icon size="18" component={SwipeLeftRound} />
            </n-el>
          );
        },
      },
    ],
  },
];

const curDataTableView = computed(() => {
  // @ts-ignore
  const rankType = tableOptions[currentTableOptions.value];
  return rankType.children?.[currentChildOption.value]?.view || rankType.view;
});

const curDataTableLabel = computed((): string | undefined => {
  const resultArr = [];

  // @ts-ignore
  const rankType = tableOptions[currentTableOptions.value];
  resultArr.push(rankType.title);

  if (rankType.children)
    resultArr.push(rankType.children[currentChildOption.value]?.title);

  resultArr.push(gameStore.dimension);

  return resultArr.join(" - ");
});
</script>

<template>
  <div class="p-4">
    <!-- 标题与功能按钮 -->
    <div class="flex justify-between items-center">
      <title-bar title="记录及排名" />

      <dropdown showDivider :options="options" :content="curDataTableLabel" />
    </div>
    <div>
      <curDataTableView />
    </div>
  </div>
</template>
