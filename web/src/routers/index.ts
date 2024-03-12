import {
  createRouter,
  createWebHashHistory,
  RouteRecordRaw,
  RouterOptions,
} from "vue-router";
import {
  MonitorRound,
  PersonSearchRound,
  EventNoteRound,
  BookmarkBorderRound,
  GridOnRound,
  KeyRound,
} from "@vicons/material";
import { markRaw } from "vue";
import { useAdminStore } from "@/store/admin";
import Home from "@/views/home/index.vue";
import Record from "@/views/record/index.vue";
import User from "@/views/user/index.vue";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    name: "Index",
    component: () => import("@/views/index/index.vue"),
    children: [
      {
        path: "/",
        name: "Home",
        component: Home,
        meta: { index: 0 },
      },
      {
        path: "/record",
        name: "Record",
        component: Record,
        meta: { index: 1 },
      },
      {
        path: "/user",
        name: "User",
        component: User,
        meta: { index: 2 },
      },
    ],
  },
  {
    path: "/edit-profile",
    name: "EditProfile",
    component: () => import("@/views/edit-profile/index.vue"),
  },
  {
    path: "/notification",
    name: "Notification",
    component: () => import("@/views/notification/index.vue"),
  },
  {
    path: "/record-detail",
    name: "RecordDetail",
    component: () => import("@/views/record-detail/index.vue"),
  },

  {
    path: "/register",
    name: "Register",
    component: () => import("@/views/register/index.vue"),
  },

  {
    path: "/login",
    name: "Login",
    component: () => import("@/views/login/index.vue"),
  },
  {
    path: "/test",
    name: "Test",
    component: () => import("@/views/test/index.vue"),
  },
  {
    path: "/admin/auth",
    name: "AdminAuth",
    component: () => import("@/views/admin/auth/index.vue"),
  },
  {
    path: "/admin",
    name: "Admin",
    component: () => import("@/views/admin/index/index.vue"),
    meta: {
      needAdminToken: true, // 需要管理员权限
    },
    children: [
      {
        path: "",
        name: "AdminHome",
        component: () => import("@/views/admin/home/index.vue"),
        meta: {
          title: "首页",
          icon: markRaw(MonitorRound),
        },
      },
      {
        path: "manage-user",
        name: "ManageUser",
        component: () => import("@/views/admin/manage-user/index.vue"),
        meta: {
          title: "用户管理",
          icon: markRaw(PersonSearchRound),
        },
      },
      {
        path: "manage-record-person",
        name: "ManageRecordPerson",
        component: () => import("@/views/admin/manage-record-person/index.vue"),
        meta: {
          rootTitle: "记录管理", // 根标题
          rootIcon: markRaw(EventNoteRound),
          title: "个人记录",
          icon: markRaw(BookmarkBorderRound),
        },
      },
      {
        path: "manage-best-single",
        name: "ManageBestSingle",
        component: () => import("@/views/admin/manage-best-single/index.vue"),
        meta: {
          rootTitle: "记录管理", // 根标题
          rootIcon: markRaw(EventNoteRound),
          title: "最佳单次",
          icon: markRaw(BookmarkBorderRound),
        },
      },
      {
        path: "manage-best-average",
        name: "ManageBestAverage",
        component: () => import("@/views/admin/manage-best-average/index.vue"),
        meta: {
          rootTitle: "记录管理", // 根标题
          rootIcon: markRaw(EventNoteRound),
          title: "最佳平均",
          icon: markRaw(BookmarkBorderRound),
        },
      },
      {
        path: "manage-best-step",
        name: "ManageBestStep",
        component: () => import("@/views/admin/manage-best-step/index.vue"),
        meta: {
          rootTitle: "记录管理", // 根标题
          rootIcon: markRaw(EventNoteRound),
          title: "最佳步数",
          icon: markRaw(BookmarkBorderRound),
        },
      },
      {
        path: "manage-scramble",
        name: "ManageScramble",
        component: () => import("@/views/admin/manage-scramble/index.vue"),
        meta: {
          title: "打乱管理",
          icon: markRaw(GridOnRound),
        },
      },
      {
        path: "setting-secret",
        name: "SettingSecret",
        component: () => import("@/views/admin/secret/index.vue"),
        meta: {
          title: "后台秘钥",
          icon: markRaw(KeyRound),
        },
      },
    ],
  },
];

const router = createRouter(<RouterOptions>{
  history: createWebHashHistory("/puzzle"),
  routes,
});

router.afterEach((to, from) => {
  // 设置页面切换动画
  if (to.meta.index !== undefined || from.meta.index !== undefined) {
    const toIndex: number = to.meta?.index as number;
    const fromIndex: number = from.meta?.index as number;

    to.meta.transition = toIndex > fromIndex ? "slide-left" : "slide-right";
  }

  // 判断是否需要管理员权限
  if (to.meta.needAdminToken) {
    const adminStore = useAdminStore();

    if (adminStore.getToken === "") {
      router.push({ name: "AdminAuth" });
    }
  }
});

export default router;
