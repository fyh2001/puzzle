<script setup lang="tsx">
import { EditNoteRound, GridOnRound } from "@vicons/material";
import { ref, onMounted } from "vue";
import { useDialog, useMessage } from "naive-ui";
import { useAdminStore } from "@/store/admin";
import { defalutAvatar } from "@/config/index";
import DefoGameMapMini from "@/components/defo-game-map-mini.vue";
import type { Pagination } from "@/types/pagination";
import type { RecordResp } from "@/types/record";

const Dialog = useDialog();
const Message = useMessage();

const adminStore = useAdminStore();

// 数据表加载状态
const isDataTableLoading = ref(false);

// 更新按钮加载状态
// const isUpdateBtnLoading = ref(false);

// 查询表单
const queryForm = ref({
  id: "",
  userId: "",
  username: "",
  nickname: "",
  dimension: null,
  recordId: "",
  stepRange: [null, null] as (number | string | null)[],
  rankRange: [null, null] as (number | string | null)[],
  breakCountRange: [null, null] as (number | string | null)[],
  dateRange: null as any,
  sorted: "asc",
  orderBy: "ranked",
  needUserInfo: true,
  needRecordDetail: true,
});

// 表格字段
const tableColumns = [
  // 排名
  {
    title: "排名",
    key: "ranked",
    width: 100,
    align: "center",
  },
  // ID
  {
    title: "ID",
    key: "id",
    width: 200,
    align: "center",
  },
  // 用户名
  {
    title: "用户名",
    key: "userInfo.username",
    width: 120,
    align: "center",
  },
  // 昵称
  {
    title: "昵称",
    key: "userInfo.nickname",
    width: 120,
    align: "center",
  },
  // 阶数
  {
    title: "阶数",
    key: "dimension",
    align: "center",
    width: 120,
    render: (row: any) => {
      return (
        <n-tag type="success" size="small">
          {row.dimension} 阶
        </n-tag>
      );
    },
  },
  // 步数
  {
    title: "步数",
    key: "recordStep",
    align: "center",
    width: 100,
  },
  // 打破记录次数
  {
    title: "打破记录次数",
    key: "recordBreakCount",
    align: "center",
    width: 100,
  },
  // 打乱
  {
    title: "打乱",
    key: "scramble",
    align: "center",
    width: 120,
    render: (row: any) => {
      return (
        <n-popover z-index={3000} display-directive="if" trigger="hover">
          {{
            trigger: () => {
              return (
                <div class="text-center">
                  <n-button
                    size="small"
                    on-click={() => {
                      handleScrambleDialogVisible(row);
                    }}
                  >
                    显示打乱
                  </n-button>
                </div>
              );
            },
            default: () => (
              <div class="w-64 h-64">
                <DefoGameMapMini
                  scramble={row.recordDetail[0].scramble}
                  dimension={row.recordDetail[0].dimension}
                  on-click={() => {
                    handleScrambleDialogVisible(row);
                  }}
                />
              </div>
            ),
          }}
        </n-popover>
      );
    },
  },
  // 创建时间
  {
    title: "创建时间",
    key: "createdAt",
    align: "center",
    width: 160,
    render: (row: any) => {
      return new Date(row.createdAt).toLocaleString();
    },
  },
  // 更新时间
  {
    title: "更新时间",
    key: "updatedAt",
    align: "center",
    width: 160,
    render: (row: any) => {
      return new Date(row.updatedAt).toLocaleString();
    },
  },
  // 操作
  {
    title: "操作",
    key: "action",
    align: "center",
    width: 200,
    render: (row: any) => {
      return (
        <div class="flex justify-center gap-2">
          <n-button
            secondary
            type="info"
            size="small"
            on-click={() => handleDetailDialog(row)}
          >
            详情
          </n-button>
        </div>
      );
    },
  },
];

// 状态选项
// const statusOptions = [
//   {
//     label: "正常",
//     value: 1,
//   },
//   {
//     label: "禁用",
//     value: 2,
//   },
//   {
//     label: "删除",
//     value: 3,
//   },
// ];

// 类型选项
const typeOptions: any = {
  1: "练习",
  2: "排行榜",
  3: "对战",
};

// 记录详情弹窗
const handleDetailDialog = (row: any) => {
  Dialog.info({
    style: { width: "60%" },
    title: "记录详情",
    icon: () => {
      return <n-icon color="#000" component={EditNoteRound} />;
    },
    content: () => {
      return (
        <div class="flex flex-col gap-4 mt-4 w-full">
          <n-descriptions
            title="最佳单次记录信息"
            label-placement="left"
            column={3}
            bordered
          >
            <n-descriptions-item label="ID">{row.id}</n-descriptions-item>
            <n-descriptions-item label="用户ID">
              {row.userId}
            </n-descriptions-item>
            <n-descriptions-item label="记录ID">
              {row.recordId}
            </n-descriptions-item>
            <n-descriptions-item label="阶数">
              {row.dimension + "阶"}
            </n-descriptions-item>
            <n-descriptions-item label="耗时">
              {row.durationFormat}
            </n-descriptions-item>
            <n-descriptions-item label="步数">
              {row.recordStep}
            </n-descriptions-item>
            <n-descriptions-item label="记录打破次数">
              {row.recordBreakCount}
            </n-descriptions-item>
            <n-descriptions-item label="排名" span={2}>
              {row.ranked}
            </n-descriptions-item>
            <n-descriptions-item label="创建时间">
              {new Date(row.createdAt).toLocaleString()}
            </n-descriptions-item>
            <n-descriptions-item label="更新时间">
              {new Date(row.createdAt).toLocaleString()}
            </n-descriptions-item>
          </n-descriptions>

          <n-descriptions title="记录列表" label-placement="left">
            <n-descriptions-item>
              <n-collapse>
                {row.recordDetail.map((data: RecordResp, index: number) => {
                  return (
                    <n-collapse-item
                      title={`第 ${index + 1} 次记录详情`}
                      name="1"
                    >
                      <n-descriptions
                        label-placement="left"
                        column={3}
                        bordered
                      >
                        <n-descriptions-item label="ID">
                          {data.id}
                        </n-descriptions-item>
                        <n-descriptions-item label="用户ID">
                          {data.userId}
                        </n-descriptions-item>
                        <n-descriptions-item label="模式">
                          {typeOptions[data.status]}
                        </n-descriptions-item>
                        <n-descriptions-item label="阶数">
                          {data.dimension + "阶"}
                        </n-descriptions-item>
                        <n-descriptions-item label="耗时">
                          {data.durationFormat}
                        </n-descriptions-item>
                        <n-descriptions-item label="步数">
                          {data.step}
                        </n-descriptions-item>
                        <n-descriptions-item label="打乱">
                          <n-ellipsis style="max-width: 240px">
                            {data.scramble}
                          </n-ellipsis>
                        </n-descriptions-item>
                        <n-descriptions-item label="还原">
                          <n-ellipsis style="max-width: 240px">
                            {data.solution}
                          </n-ellipsis>
                        </n-descriptions-item>
                        <n-descriptions-item label="idx">
                          {data.idx}
                        </n-descriptions-item>
                        <n-descriptions-item label="记录状态">
                          {
                            ({ 1: "正常", 2: "禁用", 3: "删除" } as any)[
                              data.status
                            ]
                          }
                        </n-descriptions-item>
                        <n-descriptions-item label="创建时间">
                          {new Date(data.createdAt).toLocaleString()}
                        </n-descriptions-item>
                        <n-descriptions-item label="更新时间">
                          {new Date(data.createdAt).toLocaleString()}
                        </n-descriptions-item>
                      </n-descriptions>
                    </n-collapse-item>
                  );
                })}
              </n-collapse>
            </n-descriptions-item>
          </n-descriptions>

          <n-descriptions
            title="用户信息"
            label-placement="left"
            column={3}
            bordered
          >
            <n-descriptions-item label="用户ID">
              {row.userInfo.id}
            </n-descriptions-item>
            <n-descriptions-item label="用户名">
              {row.userInfo.username}
            </n-descriptions-item>
            <n-descriptions-item label="昵称">
              {row.userInfo.nickname}
            </n-descriptions-item>
            <n-descriptions-item label="头像">
              <div class="flex items-center">
                <n-avatar size="small" src={row.avatar || defalutAvatar} />
              </div>
            </n-descriptions-item>
            <n-descriptions-item label="手机号">
              {row.userInfo.phone === "" ? "-" : row.userInfo.phone}
            </n-descriptions-item>
            <n-descriptions-item label="邮箱">
              <n-ellipsis style="max-width: 240px">
                {row.userInfo.email === "" ? "-" : row.userInfo.email}
              </n-ellipsis>
            </n-descriptions-item>
            <n-descriptions-item label="账号状态">
              {
                ({ 1: "正常", 2: "禁用", 3: "删除" } as any)[
                  row.userInfo.status
                ]
              }
            </n-descriptions-item>
            <n-descriptions-item label="创建时间">
              {new Date(row.userInfo.createdAt).toLocaleString()}
            </n-descriptions-item>
            <n-descriptions-item label="更新时间">
              {new Date(row.userInfo.createdAt).toLocaleString()}
            </n-descriptions-item>
          </n-descriptions>

          <n-descriptions title="打乱与还原" bordered>
            <n-descriptions-item
              label="打乱"
              content-style="display:flex; justify-content: center"
            >
              <div class="w-2/5 h-2/5">
                <DefoGameMapMini
                  scramble={row.recordDetail[0].scramble}
                  dimension={row.recordDetail[0].dimension}
                  solution={row.recordDetail[0].solution}
                />
              </div>
            </n-descriptions-item>
          </n-descriptions>
        </div>
      );
    },
  });
};

// 打乱弹窗
const handleScrambleDialogVisible = (row: any) => {
  Dialog.info({
    style: { width: "30%" },
    title: "打乱",
    icon: () => {
      return <n-icon component={GridOnRound} />;
    },
    content: () => {
      return (
        <div class="flex justify-center mt-4">
          <DefoGameMapMini scramble={row.scramble} dimension={row.dimension} />
        </div>
      );
    },
  });
};

// 分页
const pagination: Pagination = {
  page: 1,
  pageSize: 10,
};

// 分页改变
const handlePageChange = (page: number) => {
  pagination.page = page;
  getData();
};

// 分页大小改变
const handlePageSizeChange = (pageSize: number) => {
  pagination.pageSize = pageSize;
  getData();
};

// 获取记录
const getData = async () => {
  isDataTableLoading.value = true;

  const { code, msg } = await adminStore.fetchRecordBestStepList({
    ...queryForm.value,
    dateRange: queryForm.value.dateRange?.map((item: number) => new Date(item)),
    pagination,
  });

  if (code !== 200) {
    Message.error(msg);
  }

  isDataTableLoading.value = false;
};

// 更新记录
// const updateData = async (updateForm: RecordBestSingleReq) => {
//   isUpdateBtnLoading.value = true;
//   const { code, data, msg } = await adminStore.updateRecordList(updateForm);

//   if (code === 200) {
//     Message.success(data);
//     getData();
//   } else {
//     Message.error(msg);
//   }

//   isUpdateBtnLoading.value = false;

//   Dialog.destroyAll();
// };

// 查询表单重置
const handleQueryFormChange = () => {
  queryForm.value = {
    id: "",
    userId: "",
    username: "",
    nickname: "",
    dimension: null,
    recordId: "",
    stepRange: [null, null] as (number | string | null)[],
    rankRange: [null, null] as (number | string | null)[],
    breakCountRange: [null, null] as (number | string | null)[],
    dateRange: null as any,
    sorted: "desc",
    orderBy: "",
    needUserInfo: true,
    needRecordDetail: true,
  };
};

onMounted(() => {
  getData();
});
</script>

<template>
  <div class="p-4 w-full h-full bg-white rounded-md">
    <!-- 查找 -->
    <div class="p-4 bg-white rounded">
      <n-form label-placement="left" label-width="auto" label-align="left">
        <div class="grid grid-cols-4 gap-col-4">
          <n-form-item label="Id">
            <n-input
              v-model:value="queryForm.id"
              placeholder="请输入ID"
              clearable
            />
          </n-form-item>
          <n-form-item label="用户ID">
            <n-input
              v-model:value="queryForm.userId"
              placeholder="请输入用户ID"
              clearable
            />
          </n-form-item>
          <n-form-item label="用户名">
            <n-input
              v-model:value="queryForm.username"
              placeholder="请输入用户名"
              clearable
            />
          </n-form-item>
          <n-form-item label="用户昵称">
            <n-input
              v-model:value="queryForm.nickname"
              placeholder="请输入用户昵称"
              clearable
            />
          </n-form-item>
          <n-form-item label="记录ID">
            <n-input
              v-model:value="queryForm.recordId"
              placeholder="请输入记录ID"
              clearable
            />
          </n-form-item>
          <n-form-item label="阶数">
            <n-input-number
              class="w-full"
              v-model:value="queryForm.dimension"
              placeholder="请选择阶数"
              clearable
            />
          </n-form-item>
          <n-form-item label="步数范围">
            <n-input-number
              v-model:value="queryForm.stepRange![0]"
              :min="0"
              placeholder="起始步数"
              clearable
            />
            <span class="mx-2">至</span>
            <n-input-number
              v-model:value="queryForm.stepRange![1]"
              :min="0"
              placeholder="结束步数"
              clearable
            />
          </n-form-item>
          <n-form-item label="排名范围">
            <n-input-number
              v-model:value="queryForm.rankRange![0]"
              :min="0"
              placeholder="起始排名"
              clearable
            />
            <span class="mx-2">至</span>
            <n-input-number
              v-model:value="queryForm.rankRange![1]"
              :min="0"
              placeholder="结束排名"
              clearable
            />
          </n-form-item>
          <n-form-item label="打破记录次数范围">
            <n-input-number
              v-model:value="queryForm.breakCountRange![0]"
              :min="0"
              placeholder="起始次数"
              type="number"
              clearable
            />
            <span class="mx-2">至</span>
            <n-input-number
              v-model:value="queryForm.breakCountRange![1]"
              :min="0"
              placeholder="结束次数"
              clearable
            />
          </n-form-item>
          <n-form-item label="日期范围">
            <n-date-picker
              class="w-full"
              v-model:value="queryForm.dateRange"
              type="datetimerange"
              start-placeholder="开始时间"
              end-placeholder="结束时间"
              clearable
            />
          </n-form-item>
          <n-form-item label="排序方式">
            <n-select
              v-model:value="queryForm.orderBy"
              :options="[
                { label: 'ID', value: 'id' },
                { label: '阶数', value: 'dimension' },
                { label: '步数', value: 'record_step' },
                { label: '排名', value: 'ranked' },
                { label: '打破记录次数', value: 'record_break_count' },
                { label: '创建时间', value: 'created_at' },
                { label: '更新时间', value: 'updated_at' },
              ]"
              placeholder="排序依据"
              clearable
            />
            <n-select
              v-model:value="queryForm.sorted"
              :options="[
                { label: '升序排序', value: 'asc' },
                { label: '降序排序', value: 'desc' },
              ]"
              placeholder="排序方式"
              clearable
            />
          </n-form-item>
          <div class="flex items-center justify-end col-span-4 gap-4">
            <n-button
              class="w-20"
              strong
              secondary
              type="tertiary"
              @click="handleQueryFormChange"
            >
              重置
            </n-button>
            <n-button
              class="w-20"
              strong
              secondary
              type="primary"
              @click="getData"
            >
              查找
            </n-button>
          </div>
        </div>
      </n-form>
    </div>
    <!-- 表格 -->
    <div>
      <n-data-table
        class="mt-5"
        :loading="isDataTableLoading"
        :columns="tableColumns"
        :data="adminStore.getBestStepRecordList"
        :bordered="false"
      />
      <n-pagination
        class="mt-5 justify-center"
        show-quick-jumper
        show-size-picker
        :display-order="['size-picker', 'pages', 'quick-jumper']"
        :page="pagination.page"
        :page-size="pagination.pageSize"
        :page-sizes="[10, 15, 20, 30, 40]"
        :item-count="adminStore.getBestStepRecordTotal"
        :page-slot="3"
        @update:page="handlePageChange"
        @update:pageSize="handlePageSizeChange"
      />
    </div>
  </div>
</template>
