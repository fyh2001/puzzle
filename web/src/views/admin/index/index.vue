<script lang="tsx" setup>
import { computed, ref } from "vue";
import {
  MonitorRound,
  PersonSearchRound,
  EventNoteRound,
  BookmarkBorderRound,
  GridOnRound,
  SettingsRound,
  MenuRound,
  KeyRound,
} from "@vicons/material";
import { PersonCircleOutline } from "@vicons/ionicons5";
import router from "@/routers";
import type { MenuOption } from "naive-ui";

// 菜单是否折叠
const collapsed = ref(false);

// 菜单选项
const menuOptions: MenuOption[] = [
  // 首页
  {
    label: () => {
      return (
        <router-link to="/admin" replace>
          首页
        </router-link>
      );
    },
    breadcrumbTitle: "首页",
    key: "",
    icon: () => {
      return <n-icon size="24" component={MonitorRound} />;
    },
  },
  // 用户管理
  {
    label: () => {
      return (
        <router-link to="/admin/manage-user" replace>
          用户管理
        </router-link>
      );
    },
    breadcrumbTitle: "用户管理",
    key: "manage-user",
    icon: () => {
      return <n-icon size="24" component={PersonSearchRound} />;
    },
  },
  // 记录管理
  {
    label: "记录管理",
    breadcrumbTitle: "记录管理",
    key: "record",
    icon: () => {
      return <n-icon size="24" component={EventNoteRound} />;
    },
    children: [
      {
        label: () => {
          return (
            <router-link to="/admin/manage-record-person" replace>
              个人记录
            </router-link>
          );
        },
        breadcrumbTitle: "个人记录",
        key: "manage-record-person",
        icon: () => {
          return <n-icon size="24" component={BookmarkBorderRound} />;
        },
      },
      {
        label: () => {
          return (
            <router-link to="/admin/manage-best-single" replace>
              最佳单次
            </router-link>
          );
        },
        breadcrumbTitle: "最佳单次",
        key: "manage-best-single",
        icon: () => {
          return <n-icon size="24" component={BookmarkBorderRound} />;
        },
      },
      {
        label: () => {
          return (
            <router-link to="/admin/manage-best-average" replace>
              最佳平均
            </router-link>
          );
        },
        breadcrumbTitle: "最佳平均",
        key: "manage-best-average",
        icon: () => {
          return <n-icon size="24" component={BookmarkBorderRound} />;
        },
      },
      {
        label: () => {
          return (
            <router-link to="/admin/manage-best-step" replace>
              最佳步数
            </router-link>
          );
        },
        breadcrumbTitle: "最佳步数",
        key: "manage-best-step",
        icon: () => {
          return <n-icon size="24" component={BookmarkBorderRound} />;
        },
      },
    ],
  },
  // 打乱管理
  {
    label: () => {
      return (
        <router-link to="/admin/manage-scramble" replace>
          打乱管理
        </router-link>
      );
    },
    breadcrumbTitle: "打乱管理",
    key: "manage-scramble",
    icon: () => {
      return <n-icon size="24" component={GridOnRound} />;
    },
  },
  // 系统设置
  {
    label: "系统设置",
    breadcrumbTitle: "系统设置",
    key: "setting",
    icon: () => {
      return <n-icon size="24" component={SettingsRound} />;
    },
    children: [
      {
        label: () => {
          return (
            <router-link to="/admin/setting-secret" replace>
              后台秘钥
            </router-link>
          );
        },
        key: "setting-secret",
        icon: () => {
          return <n-icon size="24" component={KeyRound} />;
        },
      },
    ],
  },
];

// 下拉菜单选项
const dropdownOptions = [
  {
    label: "退出登录",
    key: "logout",
  },
];

// 当前面包屑数组
const currentBreadcrumbArr = computed(() => {
  const { meta } = router.currentRoute.value;
  const arr = [];

  if (meta?.rootTitle !== undefined) {
    arr.push({
      breadcrumbTitle: meta.rootTitle,
      icon: meta.rootIcon,
    });
  }

  arr.push({
    breadcrumbTitle: meta.title,
    icon: meta.icon,
  });

  return arr;
});
</script>

<template>
  <div class="w-screen h-screen">
    <n-layout class="h-screen" has-sider>
      <n-layout-sider
        class="shadow-xl"
        collapse-mode="width"
        :collapsed-width="64"
        :width="240"
        :collapsed="collapsed"
      >
        <div
          class="flex justify-center items-center gap-2 p-4 text-center text-5 text-green-7 font-bold font-mono"
        >
          <n-icon size="24" :component="GridOnRound" />
          <div v-if="!collapsed">Puzzle 管理系统</div>
        </div>
        <n-menu
          :value="router.currentRoute.value.path.replace(/^\/admin\/?/, '')"
          :collapsed="collapsed"
          :collapsed-width="64"
          :options="menuOptions"
        />
      </n-layout-sider>
      <n-layout content-class="flex flex-col h-full">
        <n-layout-header>
          <div class="flex justify-between items-center py-3 px-2">
            <div class="flex justify-between items-center gap-4">
              <n-button quaternary @click="collapsed = !collapsed">
                <n-icon size="24" :component="MenuRound" />
              </n-button>
              <n-breadcrumb>
                <n-breadcrumb-item
                  v-for="(item, index) in currentBreadcrumbArr"
                  :key="index"
                >
                  <div class="flex items-center">
                    <n-icon size="24" :component="item.icon" />
                    <div class="ml-2 font-mono">{{ item.breadcrumbTitle }}</div>
                  </div>
                </n-breadcrumb-item>
              </n-breadcrumb>
            </div>
            <div>
              <n-dropdown trigger="click" :options="dropdownOptions">
                <n-button quaternary>
                  <n-icon size="24" :component="PersonCircleOutline" />
                  <div class="ml-2 text-4 font-bold font-mono">Admin</div>
                </n-button>
              </n-dropdown>
            </div>
          </div>
        </n-layout-header>
        <n-layout-content class="flex-grow p-4">
          <router-view v-slot="{ Component, route }">
            <transition
              :name="route.meta.transition as string || 'fade'"
              mode="out-in"
            >
              <component :is="Component" />
            </transition>
          </router-view>
        </n-layout-content>
      </n-layout>
    </n-layout>
  </div>
</template>
