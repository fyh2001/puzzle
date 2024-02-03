import {
  createRouter,
  createWebHashHistory,
  RouteRecordRaw,
  RouterOptions,
} from "vue-router";

import Index from "@/views/index/index.vue";
import Home from "@/views/home/index.vue";
import Record from "@/views/record/index.vue";
import User from "@/views/user/index.vue";
import Register from "@/views/register/index.vue";
import Login from "@/views/login/index.vue";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    name: "Index",
    component: Index,
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
    path: "/register",
    name: "Register",
    component: Register,
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
];

const router = createRouter(<RouterOptions>{
  history: createWebHashHistory("/puzzle"),
  routes,
});

router.afterEach((to, from) => {
  const toIndex: number = to.meta.index as number;
  const fromIndex: number = from.meta.index as number;

  to.meta.transition = toIndex > fromIndex ? "slide-left" : "slide-right";
});

export default router;
