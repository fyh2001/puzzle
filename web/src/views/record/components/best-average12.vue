<script lang="tsx" setup>
import { ref, onMounted, watch } from "vue";
import { useRecordStore } from "@/store/record";
import { useGameStore } from "@/store/game";
import type { Pagination } from "@/types/pagination";

const recordStore = useRecordStore();
const gameStore = useGameStore();

const isLoading = ref(false);

// 表格字段
const tableColumns = [
  {
    title: "No",
    dataIndex: "index",
    key: "index",
    width: "65",
    render: (_: Object, index: number) => {
      return (
        recordStore.rankedRecord.bestAverage12.total -
        (pagination.page - 1) * 10 -
        index
      );
    },
  },
  {
    title: "用户",
    dataIndex: "nickname",
    key: "nickname",
  },
  {
    title: "平均时长",
    dataIndex: "duration",
    key: "duration",
    ellipsis: {
      tooltip: true,
    },
  },
];

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
  await recordStore.getBestAverage12Records({
    pagination,
    dimension: gameStore.getDimension,
  });

  isLoading.value = false;
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
});
</script>

<template>
  <div>
    <n-data-table
      class="mt-5"
      :loading="isLoading"
      :columns="tableColumns"
      :data="recordStore.getBestAverage12RecordList"
      :bordered="false"
    />
    <n-pagination
      class="mt-5 justify-end"
      :page="pagination.page"
      :page-size="pagination.pageSize"
      :item-count="recordStore.rankedRecord.bestAverage12.total"
      :page-slot="3"
      @update:page="handlePageChange"
    />
  </div>
</template>

<style>
.n-data-table-th__title {
  text-align: center;
}

td.n-data-table-td.n-data-table-td {
  text-align: center;
}

/**
.n-data-table .n-data-table-td.n-data-table-td--summary {
  @apply bg-indigo-50 border border-indigo-200;
  background-color: v-bind(--rank-summary-background-color);
}
*/
</style>
