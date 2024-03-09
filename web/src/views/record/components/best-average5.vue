<script lang="tsx" setup>
import { ref, onMounted, watch, computed } from "vue";
import { useRecordStore } from "@/store/record";
import { useGameStore } from "@/store/game";
import { useUserStore } from "@/store/user";
import { useThemeStore } from "@/store/theme";
import { recordRequest } from "api/record";
import { formatDurationInRecord } from "@/utils/time";
import type { Pagination } from "@/types/pagination";
import type { RecordBestAverageResp } from "@/types/record";
import router from "@/routers";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const recordStore = useRecordStore();
const gameStore = useGameStore();
const userStore = useUserStore();
const themeStore = useThemeStore();

const isLoading = ref(false);

// 总结栏背景颜色
const summaryBackgroundColor = computed(
  () =>
    (
      themeStore.loadedThemesMap.get(themeStore.themeColor)?.[
        themeStore.getDarkMode ? "dark" : "light"
      ] as any
    )?.common?.rankSummaryBackgroundColor
);

// 总结栏文字颜色
const summaryTextColor = computed(
  () =>
    (
      themeStore.loadedThemesMap.get(themeStore.themeColor)?.[
        themeStore.getDarkMode ? "dark" : "light"
      ] as any
    )?.common?.rankSummaryTextColor
);

// 个人记录
const recordPersonData = ref<RecordBestAverageResp>();

// 表格字段
const tableColumns = computed(() => [
  {
    title: "No",
    dataIndex: "ranked",
    key: "ranked",
    width: "65",
  },
  {
    title: t("record.table.column.user"),
    dataIndex: "nickname",
    key: "nickname",
  },
  {
    title: t("record.table.column.duration"),
    dataIndex: "duration",
    key: "duration",
    ellipsis: {
      tooltip: true,
    },
  },
]);

// 总结栏
const tableSummary = () => {
  return {
    ranked: {
      value: <div>{recordPersonData.value?.ranked || "-"}</div>,
      colSpan: 1,
    },
    nickname: {
      value: <div>{recordPersonData.value?.userInfo?.nickname || "-"}</div>,
      colSpan: 1,
    },
    duration: {
      value: (
        <div>{recordPersonData.value?.recordAverageDurationFormat || "-"}</div>
      ),
      colSpan: 1,
    },
  };
};

// 表格行点击
const handleTableRowClick = (rowData: any) => {
  return {
    onClick: () => {
      router.push({
        name: "RecordDetail",
        query: {
          data: JSON.stringify(rowData),
          type: "bestAverage",
        },
      });
    },
  };
};

// 分页
const pagination: Pagination = {
  page: 1,
  pageSize: 10,
};

// 分页改变
const handlePageChange = (page: number) => {
  pagination.page = page;
  getRecords();
};

// 获取记录
const getRecords = async () => {
  isLoading.value = true;
  await recordStore.getBestAverage5Records({
    pagination,
    dimension: gameStore.getDimension,
    orderBy: "ranked",
    needUserInfo: true,
  });

  isLoading.value = false;
};

// 获取个人记录
const getRecordPersion = async () => {
  const {
    data: { code, data: recordPersonResp },
  } = await recordRequest.listBestAverage({
    pagination: {
      page: 1,
      pageSize: 1,
    },
    type: 5,
    dimension: gameStore.getDimension,
    userId: userStore.getUser.id,
    needUserInfo: true,
  });

  if (code === 200) {
    recordPersonData.value = recordPersonResp.records[0] || {};
    if (recordPersonData.value.recordAverageDuration) {
      recordPersonData.value.recordAverageDurationFormat =
        formatDurationInRecord(recordPersonData.value.recordAverageDuration);
    }
  }
};

// 监听阶数和模式变化
watch(
  () => [gameStore.getDimension, gameStore.getGameMode],
  () => {
    // 重新获取记录
    getRecords();
  }
);

onMounted(() => {
  getRecords();
  getRecordPersion();
});
</script>

<template>
  <div>
    <n-data-table
      class="mt-5"
      :loading="isLoading"
      :columns="tableColumns"
      :data="recordStore.getBestAverage5RecordList"
      :row-props="handleTableRowClick"
      :summary="tableSummary"
      :bordered="false"
    />
    <n-pagination
      class="mt-5 justify-end"
      :page="pagination.page"
      :page-size="pagination.pageSize"
      :item-count="recordStore.rankedRecord.bestAverage5.total"
      :page-slot="3"
      @update:page="handlePageChange"
    />
  </div>
</template>

<style scoped>
:deep .n-data-table-th__title {
  text-align: center;
}

:deep td.n-data-table-td.n-data-table-td {
  text-align: center;
}

:deep .n-data-table .n-data-table-td.n-data-table-td--summary {
  background-color: v-bind(summaryBackgroundColor);
  color: v-bind(summaryTextColor);
}

:deep .n-data-table .n-data-table-wrapper {
  overflow: hidden;
  border-radius: 0.5rem;
}
</style>
