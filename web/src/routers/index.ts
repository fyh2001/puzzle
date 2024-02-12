import {
    createRouter,
    createWebHashHistory,
    RouteRecordRaw,
    RouterOptions,
} from "vue-router";

const routes: RouteRecordRaw[] = [
    {
        path: "/",
        name: "Index",
        component: () => import("@/views/index/index.vue"),
        children: [
            {
                path: "/",
                name: "Home",
                component: () => import("@/views/home/index.vue"),
                meta: {index: 0},
            },
            {
                path: "/record",
                name: "Record",
                component: () => import("@/views/record/index.vue"),
                meta: {index: 1},
            },
            {
                path: "/user",
                name: "User",
                component: () => import("@/views/user/index.vue"),
                meta: {index: 2},
            },
        ],
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
        children: [
            {
                path: "",
                name: "AdminHome",
                component: () => import("@/views/admin/home/index.vue"),
            },
            {
                path: "manage-user",
                name: "ManageUser",
                component: () => import("@/views/admin/manage-user/index.vue"),
            },
            {
                path: "manage-record-person",
                name: "ManageRecordPerson",
                component: () => import("@/views/admin/manage-record-person/index.vue"),
            },
            {
                path: "manage-best-single",
                name: "ManageBestSingle",
                component: () => import("@/views/admin/manage-best-single/index.vue"),
            },
            {
                path: "manage-best-average",
                name: "ManageBestAverage",
                component: () => import("@/views/admin/manage-best-average/index.vue"),
            },
            {
                path: "manage-best-step",
                name: "ManageBestStep",
                component: () => import("@/views/admin/manage-best-step/index.vue"),
            },
            {
                path: "manage-scramble",
                name: "ManageScramble",
                component: () => import("@/views/admin/manage-scramble/index.vue"),
            },
            {
                path: "setting",
                name: "Setting",
                component: () => import("@/views/admin/setting/index.vue"),
            }
        ]
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
