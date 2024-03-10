<script lang="tsx" setup>
import { ref, onMounted, watch, computed } from "vue";
import router from "@/routers";
import { useRecordStore } from "@/store/record";
import { useGameStore } from "@/store/game";
import { useUserStore } from "@/store/user";
import type { Pagination } from "@/types/pagination";
import { useI18n } from "vue-i18n";

const { t } = useI18n();

const recordStore = useRecordStore();
const gameStore = useGameStore();
const userStore = useUserStore();

const isLoading = ref(false);

// 表格字段
const tableColumns = computed(() => [
  {
    title: "No",
    dataIndex: "index",
    key: "index",
    width: "65",
    render: (_: Object, index: number) => {
      return (
        recordStore.personRecord.total - (pagination.page - 1) * 10 - index
      );
    },
  },
  {
    title: t("record.table.column.duration"),
    dataIndex: "duration",
    key: "duration",
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: t("record.table.column.step"),
    dataIndex: "step",
    key: "step",
    width: "80",
  },
  {
    title: "Ao5",
    dataIndex: "ao5",
    key: "ao5",
    ellipsis: {
      tooltip: true,
    },
  },
  {
    title: "Ao12",
    dataIndex: "ao12",
    key: "ao12",
    ellipsis: {
      tooltip: true,
    },
  },
]);

// 表格行点击
const handleTableRowClick = (rowData: any) => {
  return {
    onClick: () => {
      router.push({
        name: "RecordDetail",
        query: {
          data: JSON.stringify(rowData),
          type: "person",
        },
      });
    },
  };
};

// 分页
const pagination: Pagination = {
  page: 1,
  pageSize: 19,
  offset: 0,
};

// 分页改变
const handlePageChange = (page: number) => {
  pagination.page = page;
  pagination.offset = (pagination.page - 1) * 10;
  getRecords();
};

// 获取记录
const getRecords = async () => {
  isLoading.value = true;
  await recordStore.getPersonRecords({
    pagination,
    dimension: gameStore.getDimension,
    type: gameStore.getGameMode,
    userId: userStore.getUser.id,
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
      :data="recordStore.getPersonRecordList"
      :row-props="handleTableRowClick"
      :bordered="false"
    />
    <n-pagination
      class="mt-5 justify-end"
      :page="pagination.page"
      :page-size="10"
      :item-count="recordStore.personRecord.total"
      :page-slot="3"
      @update:page="handlePageChange"
    />
  </div>
</template>

<style scoped>
:deep(.n-data-table-th__title) {
  text-align: center;
}

:deep(td.n-data-table-td.n-data-table-td) {
  text-align: center;
}

:deep(.n-data-table .n-data-table-wrapper) {
  overflow: hidden;
  border-radius: 0.5rem;
}
</style>
