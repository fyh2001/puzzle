<script lang="tsx" setup>
import { ref, computed, markRaw, onMounted, Component } from "vue";
import {
  MovingRound,
  CableRound,
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
import bestSingle from "@/views/record/components/best-single.vue";
import bestAverage5 from "@/views/record/components/best-average5.vue";
import bestAverage12 from "@/views/record/components/best-average12.vue";
import bestStep from "@/views/record/components/best-step.vue";
import { useGameStore } from "@/store/game";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const gameStore = useGameStore();

// 自定义表格选项
const tableOptions = computed(() => ({
  mine: {
    title: t("record.dropdown.mine.label"),
    view: recordPerson,
    disabled: false,
    children: {
      practice: {
        title: t("record.dropdown.mine.content.practice"),
        icon: markRaw(WorkspacePremiumRound),
        disabled: false,
        methods: () => {
          gameStore.setGameMode(1);
        },
      },
      myRank: {
        title: t("record.dropdown.mine.content.rank"),
        icon: markRaw(MovingRound),
        disabled: false,
        methods: () => {
          gameStore.setGameMode(2);
        },
      },
      myBattle: {
        title: t("record.dropdown.mine.content.battle"),
        icon: markRaw(CableRound),
        disabled: true,
        methods: () => {
          gameStore.setGameMode(3);
        },
      },
    },
  },
  all: {
    title: t("record.dropdown.rank.label"),
    disabled: false,
    children: {
      bestSingle: {
        title: t("record.dropdown.rank.content.bestSingle"),
        icon: markRaw(TimerOutlined),
        disabled: false,
        view: bestSingle,
      },
      bestAverage5: {
        title: t("record.dropdown.rank.content.bestAverage5"),
        icon: markRaw(Filter5Round),
        disabled: false,
        view: bestAverage5,
      },
      bestAverage12: {
        title: t("record.dropdown.rank.content.bestAverage12"),
        icon: markRaw(Filter9PlusRound),
        disabled: false,
        view: bestAverage12,
      },
      bestStep: {
        title: t("record.dropdown.rank.content.bestStep"),
        icon: markRaw(SwipeLeftRound),
        disabled: false,
        view: bestStep,
      },
    },
  },
  weekly: {
    title: t("record.dropdown.rankWeekly.label"),
    disabled: true,
    children: {
      bestSingle: {
        title: t("record.dropdown.rankWeekly.content.bestSingle"),
        icon: markRaw(TimerOutlined),
        disabled: true,
      },
      bestAverage5: {
        title: t("record.dropdown.rankWeekly.content.bestAverage5"),
        icon: markRaw(Filter5Round),
        disabled: true,
      },
      bestAverage12: {
        title: t("record.dropdown.rankWeekly.content.bestAverage12"),
        icon: markRaw(Filter9PlusRound),
        disabled: true,
      },
      bestStepCount: {
        title: t("record.dropdown.rankWeekly.content.bestStep"),
        icon: markRaw(SwipeLeftRound),
        disabled: true,
      },
    },
  },
}));

const IIcon = (icon: Component) => () =>
  (
    <n-el class="flex items-center" style="color: var(--primary-color)">
      <n-icon size="18" component={icon} />
    </n-el>
  );

// 功能下拉框选项
const options = computed(() => [
  // 阶数
  {
    label: t("record.dropdown.dimension.label"),
    key: "dimension",
    icon: IIcon(DashboardCustomizeRound),
    children: [
      {
        label: "3x3",
        key: "3",
        icon: IIcon(Looks3Outlined),
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
    label: t("record.dropdown.mine.label"),
    key: "record",
    disabled: tableOptions.value.mine.disabled,
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={RocketLaunchOutlined} />
        </n-el>
      );
    },
    children: Object.entries(tableOptions.value.mine.children).map(
      ([key, val]) => ({
        label: val.title,
        key: key,
        disabled: val?.disabled || false,
        props: {
          onClick: () => {
            currentTableOptions.value = "mine";
            currentChildOption.value = key;
            val.methods();
          },
        },
        icon: () => (
          <n-el class="flex items-center" style="color: var(--primary-color)">
            <n-icon size="18" component={val.icon} />
          </n-el>
        ),
      })
    ),
  },
  // 总排名
  {
    label: t("record.dropdown.rank.label"),
    key: "rank",
    disabled: tableOptions.value.all?.disabled,
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={WorkspacePremiumRound} />
        </n-el>
      );
    },
    children: Object.entries(tableOptions.value.all.children).map(
      ([key, val]) => ({
        label: val.title,
        key: key,
        disabled: val?.disabled || false,
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
      })
    ),
  },
  // 周排名
  {
    label: t("record.dropdown.rankWeekly.label"),
    key: "weekRank",
    disabled: true,
    icon: () => {
      return (
        <n-el class="flex items-center" style="color: var(--primary-color)">
          <n-icon size="18" component={CalendarMonthRound} />
        </n-el>
      );
    },
    children: Object.entries(tableOptions.value.weekly.children).map(
      ([key, val]) => ({
        label: val.title,
        key: key,
        disabled: val?.disabled || false,
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
      })
    ),
  },
]);

// 当前表格选项
const currentTableOptions = ref("mine");
// 当前子表格选项
const currentChildOption = ref("");

// 当前表格视图
const curDataTableView = computed(() => {
  // @ts-ignore
  const rankType = tableOptions.value[currentTableOptions.value];
  return rankType.children?.[currentChildOption.value]?.view || rankType.view;
});

// 当前表格标签
const curDataTableLabel = computed((): string | undefined => {
  const resultArr = [];

  // @ts-ignore
  // 主表格标题
  const rankType = tableOptions.value[currentTableOptions.value];
  resultArr.push(rankType.title);

  // 子表格标题
  if (rankType.children)
    resultArr.push(rankType.children[currentChildOption.value]?.title);

  // 阶数
  resultArr.push(gameStore.dimension + "x" + gameStore.dimension);

  return resultArr.join(" - ");
});

onMounted(() => {
  const modeToChildOptionMap: Record<number, string> = {
    1: "practice",
    2: "myRank",
    3: "myBattle",
  };

  currentChildOption.value = modeToChildOptionMap[gameStore.getGameMode];
});
</script>

<template>
  <div class="p-4">
    <!-- 标题与功能按钮 -->
    <div class="flex justify-between items-center">
      <title-bar
        :title="t('record.title')"
        :textSize="[
          {
            size: '17px',
            language: 'us',
          },
          {
            size: '17px',
            language: 'jp',
          },
        ]"
      />

      <dropdown showDivider :options="options" :content="curDataTableLabel" />
    </div>
    <div>
      <curDataTableView />
    </div>
  </div>
</template>
