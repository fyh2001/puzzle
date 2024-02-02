import {
  createRouter,
  createWebHashHistory,
  RouteRecordRaw,
  RouterOptions,
} from "vue-router";

import Index from "@/views/index/index.vue";
import Home from "@/views/home/index.vue";
import Register from "@/views/register/index.vue";

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
    ],
  },
  {
    path: "/register",
    name: "Register",
    component: Register,
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
