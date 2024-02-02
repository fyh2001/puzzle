import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { resolve } from "path";
import UnoCSS from "unocss/vite";

import presetIcons from "@unocss/preset-icons";
import transformerDirective from "@unocss/transformer-directives";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { NaiveUiResolver } from "unplugin-vue-components/resolvers";
// https://vitejs.dev/config/
export default defineConfig({
  base: "/puzzle",
  resolve: {
    alias: {
      "@": resolve(__dirname, "./src"),
      api: resolve(__dirname, "./src/api"),
    },
  },

  plugins: [
    vue(),
    UnoCSS({
      transformers: [transformerDirective()],
      presets: [presetIcons({})],
    }),
    AutoImport({
      imports: [
        "vue",
        {
          "naive-ui": [
            "useDialog",
            "useMessage",
            "useNotification",
            "useLoadingBar",
          ],
        },
      ],
      resolvers: [NaiveUiResolver()],
    }),
    Components({
      resolvers: [NaiveUiResolver()],
    }),
  ],
});