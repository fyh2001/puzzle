<script setup lang="tsx">
import { EditNoteRound } from "@vicons/material";
import { ref, onMounted } from "vue";
import { useDialog, useMessage } from "naive-ui";
import { useAdminStore } from "@/store/admin";
import { defaultAvatar } from "@/config/index";
import { md5 } from "js-md5";
import { salt } from "@/config/index";
import type { Pagination } from "@/types/pagination";
import type { UserReq } from "@/types/user";
const Dialog = useDialog();
const Message = useMessage();

const adminStore = useAdminStore();

// 数据表加载状态
const isDataTableLoading = ref(false);

// 更新按钮加载状态
const isUpdateBtnLoading = ref(false);

// 查询表单
const queryForm = ref({
  id: "",
  username: "",
  nickname: "",
  phone: "",
  email: "",
  status: null,
  dateRange: null as any,
});

// 表格字段
const tableColumns = [
  {
    title: "ID",
    key: "id",
    align: "center",
  },
  {
    title: "头像",
    key: "avatar",
    align: "center",
    render: (row: any) => {
      return (
        <div class="flex justify-center items-center">
          <n-avatar size="small" src={row.avatar || defaultAvatar} />
        </div>
      );
    },
  },
  {
    title: "用户名",
    key: "username",
    align: "center",
  },
  {
    title: "昵称",
    key: "nickname",
    align: "center",
  },

  {
    title: "手机",
    key: "phone",
    align: "center",
  },
  {
    title: "邮箱",
    key: "email",
    align: "center",
  },
  {
    title: "账号状态",
    key: "status",
    align: "center",
    render: (row: any) => {
      const statusMap: any = {
        1: "正常",
        2: "禁用",
        3: "删除",
      };
      // return statusMap[row.status] || "";

      return (
        <n-tag
          type={
            row.status === 1
              ? "success"
              : row.status === 2
              ? "warning"
              : "error"
          }
          size="small"
        >
          {statusMap[row.status] || ""}
        </n-tag>
      );
    },
  },
  {
    title: "创建时间",
    key: "createdAt",
    align: "center",
    render: (row: any) => {
      return new Date(row.createdAt).toLocaleString();
    },
  },
  {
    title: "更新时间",
    key: "updatedAt",
    align: "center",
    render: (row: any) => {
      return new Date(row.updatedAt).toLocaleString();
    },
  },
  {
    title: "操作",
    key: "action",
    align: "center",
    render: (row: any) => {
      return (
        <div>
          <n-button
            secondary
            type="primary"
            size="small"
            on-click={() => handleEditDialog(row)}
          >
            编辑
          </n-button>
        </div>
      );
    },
  },
];

// 状态选项
const statusOptions = [
  {
    label: "正常",
    value: 1,
  },
  {
    label: "禁用",
    value: 2,
  },
  {
    label: "删除",
    value: 3,
  },
];

// 编辑用户
const handleEditDialog = (row: any) => {
  const userInfo = ref({
    ...row,
    password: "",
  });

  Dialog.info({
    style: { width: "50%" },
    title: "编辑用户",
    icon: () => {
      return <n-icon color="#000" component={EditNoteRound} />;
    },
    content: () => {
      return (
        <div class="mt-4">
          <n-form
            label-placement="left"
            label-width="auto"
            label-align="left"
            require-mark-placement="right-hanging"
          >
            <n-form-item label="ID" required>
              <n-input disabled v-model={[userInfo.value.id, "value"]} />
            </n-form-item>
            <n-form-item label="用户名" required>
              <n-input
                v-model={[userInfo.value.username, "value"]}
                placeholder="修改用户名"
              />
            </n-form-item>
            <n-form-item label="昵称" required>
              <n-input
                v-model={[userInfo.value.nickname, "value"]}
                placeholder="修改昵称"
              />
            </n-form-item>
            <n-form-item label="密码">
              <n-input
                v-model={[userInfo.value.password, "value"]}
                placeholder="重置密码"
              />
            </n-form-item>
            <n-form-item label="手机">
              <n-input
                v-model={[userInfo.value.phone, "value"]}
                placeholder="修改手机号"
              />
            </n-form-item>
            <n-form-item label="邮箱">
              <n-input
                v-model={[userInfo.value.email, "value"]}
                placeholder="修改邮箱"
              />
            </n-form-item>
            <n-form-item label="账号状态" required>
              <n-select
                v-model={[userInfo.value.status, "value"]}
                options={statusOptions}
              />
            </n-form-item>
          </n-form>
        </div>
      );
    },
    action: () => {
      return (
        <div class="flex justify-end gap-4">
          <n-button strong secondary type="tertiary" on-click={() => {}}>
            取消
          </n-button>
          <n-button
            strong
            secondary
            type="primary"
            loading={isUpdateBtnLoading.value}
            disabled={isUpdateBtnLoading.value}
            on-click={() => {
              updateData(userInfo.value);
            }}
          >
            确定
          </n-button>
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

  const { code, msg } = await adminStore.fetchUserList({
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
const updateData = async (updateForm: UserReq) => {
  isUpdateBtnLoading.value = true;

  updateForm.password =
    updateForm.password === "" ? "" : md5(updateForm.password + salt);

  const { code, data, msg } = await adminStore.updateUserList(updateForm);

  if (code === 200) {
    Message.success(data);
    getData();
  } else {
    Message.error(msg);
  }

  isUpdateBtnLoading.value = false;

  Dialog.destroyAll();
};

// 查询表单重置
const handleQueryFormChange = () => {
  queryForm.value = {
    id: "",
    username: "",
    nickname: "",
    phone: "",
    email: "",
    status: null,
    dateRange: null as any,
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
          <n-form-item label="ID">
            <n-input v-model:value="queryForm.id" placeholder="请输入用户ID" />
          </n-form-item>
          <n-form-item label="用户名">
            <n-input
              v-model:value="queryForm.username"
              placeholder="请输入用户名"
            />
          </n-form-item>
          <n-form-item label="昵称">
            <n-input
              v-model:value="queryForm.nickname"
              placeholder="请输入昵称"
            />
          </n-form-item>
          <n-form-item label="手机">
            <n-input v-model:value="queryForm.phone" placeholder="请输入手机" />
          </n-form-item>
          <n-form-item label="邮箱">
            <n-input v-model:value="queryForm.email" placeholder="请输入邮箱" />
          </n-form-item>
          <n-form-item label="状态">
            <n-select
              clearable
              v-model:value="queryForm.status"
              :options="statusOptions"
              placeholder="请选择状态"
            />
          </n-form-item>
          <n-form-item label="日期范围">
            <n-date-picker
              class="w-full"
              v-model:value="queryForm.dateRange"
              type="datetimerange"
              clearable
              start-placeholder="开始时间"
              end-placeholder="结束时间"
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
        :data="adminStore.getUserList"
        :bordered="false"
      />
      <n-pagination
        class="mt-5 justify-center"
        show-quick-jumper
        show-size-picker
        :display-order="[, 'size-picker', 'pages', 'quick-jumper']"
        :page="pagination.page"
        :page-size="pagination.pageSize"
        :page-sizes="[10, 20, 30, 40]"
        :item-count="adminStore.getUserTotal"
        :page-slot="3"
        @update:page="handlePageChange"
        @update:pageSize="handlePageSizeChange"
      />
    </div>
  </div>
</template>
