import { createApp } from "vue";
import "@arco-design/web-vue/dist/arco.css";
import "./style.css";
import App from "./App.vue";

const app = createApp(App);

// Unocss
import "@unocss/reset/tailwind.css";
import "uno.css";

// pinia全局状态管理
import { createPinia } from "pinia";
import piniaPluginPersistedstate from "pinia-plugin-persistedstate"; // pinia持久化插件
const pinia = createPinia();
pinia.use(piniaPluginPersistedstate);
app.use(pinia);

// vue-i18n国际化
import { createI18n } from "vue-i18n";
import languages from "@/languages/index.ts";
import { useCommonStore } from "./store/common";
const locale = useCommonStore().getLanguage;
const i18n = createI18n({
  legacy: false,
  locale: locale,
  fallbackLocale: "cn",
  messages: languages,
});
app.use(i18n);

//vue-router路由管理
import router from "@/routers/index.ts";
app.use(router);

app.mount("#app");
