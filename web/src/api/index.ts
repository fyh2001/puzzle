import axios from "axios";
import { useUserStore } from "@/store/user";
import { useAdminStore } from "@/store/admin";
import router from "@/routers";

// const baseURL = "http://localhost:8081/api";
// const baseURL = "http://192.168.31.141:8081/api";
// const baseURL = "http://10.211.55.3:8081/api";
const baseURL = "http://139.9.7.92:8081/api";

const request = axios.create({
  baseURL,
  timeout: 1000,
  headers: {
    "Content-Type": "application/json",
  },
});

/**
 * 请求拦截器
 */
request.interceptors.request.use(
  (config) => {
    // const Message = useMessage();

    const userStore = useUserStore();
    const adminStore = useAdminStore();

    if (config.url?.includes("admin")) {
      config.headers.Authorization = adminStore.getToken;
    } else {
      config.headers.Authorization = userStore.getToken;
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

/**
 * 响应拦截器
 */
request.interceptors.response.use(
  (response) => {
    if (response.data.msg === "无权限访问Admin") {
      const adminStore = useAdminStore();
      adminStore.resetaAuthorization();
      router.push("/admin/auth");
    }

    return response;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 暴露所写的内容
export { request, baseURL };
