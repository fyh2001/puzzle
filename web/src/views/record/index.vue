<script lang="tsx" setup>
import { ref, computed } from "vue";
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
import View from "@/views/record/components/record-person.vue";
import { useGameStore } from "@/store/game";

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
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={WorkspacePremiumRound} />
        </n-el>
      );
    },
    children: [
      {
        label: "最佳单次",
        key: "bestSingle",
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
        key: "bestAverage5",
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
        key: "bestAverage12",
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
        key: "bestStepCount",
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

const curDataTable = ref("record");

const curDataTableLabel = computed((): string | undefined => {
  if (["3", "4", "5", "6"].includes(curDataTable.value)) {
    return curDataTableLabel.value;
  }

  for (const item of options) {
    if (item.key === curDataTable.value)
      return `${item.label} - ${gameStore.dimension} 阶`;
    if (item.children) {
      const child = item.children.find(
        (child) => child.key === curDataTable.value
      );
      if (child)
        return `${item.label} - ${child.label} - ${gameStore.dimension} 阶`;
    }
  }
});
</script>

<template>
  <div class="p-4">
    <!-- 标题与功能按钮 -->
    <div class="flex justify-between items-center">
      <title-bar title="记录及排名" />

      <dropdown
        showDivider
        :options="options"
        :content="curDataTableLabel"
        @select="($event) => (curDataTable = $event)"
      />
    </div>
    <div>
      <View />
    </div>
  </div>
</template>
