// vite.config.ts
import { defineConfig } from "file:///Users/huangy/Desktop/puzzle/web/node_modules/.pnpm/vite@5.0.12_@types+node@20.11.15/node_modules/vite/dist/node/index.js";
import vue from "file:///Users/huangy/Desktop/puzzle/web/node_modules/.pnpm/@vitejs+plugin-vue@4.6.2_vite@5.0.12_vue@3.4.15/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import { resolve } from "path";
import UnoCSS from "file:///Users/huangy/Desktop/puzzle/web/node_modules/.pnpm/unocss@0.58.4_postcss@8.4.33_vite@5.0.12/node_modules/unocss/dist/vite.mjs";
import vueJsx from "file:///Users/huangy/Desktop/puzzle/web/node_modules/.pnpm/@vitejs+plugin-vue-jsx@3.1.0_vite@5.0.12_vue@3.4.15/node_modules/@vitejs/plugin-vue-jsx/dist/index.mjs";
import presetIcons from "file:///Users/huangy/Desktop/puzzle/web/node_modules/.pnpm/@unocss+preset-icons@0.58.4/node_modules/@unocss/preset-icons/dist/index.mjs";
import transformerDirective from "file:///Users/huangy/Desktop/puzzle/web/node_modules/.pnpm/@unocss+transformer-directives@0.58.4/node_modules/@unocss/transformer-directives/dist/index.mjs";
import AutoImport from "file:///Users/huangy/Desktop/puzzle/web/node_modules/.pnpm/unplugin-auto-import@0.17.5_@vueuse+core@10.9.0/node_modules/unplugin-auto-import/dist/vite.js";
import Components from "file:///Users/huangy/Desktop/puzzle/web/node_modules/.pnpm/unplugin-vue-components@0.26.0_vue@3.4.15/node_modules/unplugin-vue-components/dist/vite.js";
import { NaiveUiResolver } from "file:///Users/huangy/Desktop/puzzle/web/node_modules/.pnpm/unplugin-vue-components@0.26.0_vue@3.4.15/node_modules/unplugin-vue-components/dist/resolvers.js";
import { ArcoResolver } from "file:///Users/huangy/Desktop/puzzle/web/node_modules/.pnpm/unplugin-vue-components@0.26.0_vue@3.4.15/node_modules/unplugin-vue-components/dist/resolvers.js";
var __vite_injected_original_dirname = "/Users/huangy/Desktop/puzzle/web";
var vite_config_default = defineConfig({
  base: "/puzzle",
  resolve: {
    alias: {
      "@": resolve(__vite_injected_original_dirname, "./src"),
      api: resolve(__vite_injected_original_dirname, "./src/api/methods")
    }
  },
  plugins: [
    vue({
      template: {
        compilerOptions: {
          isCustomElement: (tag) => tag.startsWith("defo-")
        }
      }
    }),
    vueJsx(),
    UnoCSS({
      transformers: [transformerDirective()],
      presets: [presetIcons({})]
    }),
    AutoImport({
      imports: [
        "vue",
        {
          "naive-ui": [
            "useDialog",
            "useMessage",
            "useNotification",
            "useLoadingBar"
          ]
        }
      ],
      resolvers: [NaiveUiResolver(), ArcoResolver()]
    }),
    Components({
      resolvers: [
        NaiveUiResolver(),
        ArcoResolver({
          sideEffect: true
        })
      ]
    })
  ]
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCIvVXNlcnMvaHVhbmd5L0Rlc2t0b3AvcHV6emxlL3dlYlwiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9maWxlbmFtZSA9IFwiL1VzZXJzL2h1YW5neS9EZXNrdG9wL3B1enpsZS93ZWIvdml0ZS5jb25maWcudHNcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfaW1wb3J0X21ldGFfdXJsID0gXCJmaWxlOi8vL1VzZXJzL2h1YW5neS9EZXNrdG9wL3B1enpsZS93ZWIvdml0ZS5jb25maWcudHNcIjtpbXBvcnQgeyBkZWZpbmVDb25maWcgfSBmcm9tIFwidml0ZVwiO1xuaW1wb3J0IHZ1ZSBmcm9tIFwiQHZpdGVqcy9wbHVnaW4tdnVlXCI7XG5pbXBvcnQgeyByZXNvbHZlIH0gZnJvbSBcInBhdGhcIjtcbmltcG9ydCBVbm9DU1MgZnJvbSBcInVub2Nzcy92aXRlXCI7XG5pbXBvcnQgdnVlSnN4IGZyb20gXCJAdml0ZWpzL3BsdWdpbi12dWUtanN4XCI7XG5pbXBvcnQgcHJlc2V0SWNvbnMgZnJvbSBcIkB1bm9jc3MvcHJlc2V0LWljb25zXCI7XG5pbXBvcnQgdHJhbnNmb3JtZXJEaXJlY3RpdmUgZnJvbSBcIkB1bm9jc3MvdHJhbnNmb3JtZXItZGlyZWN0aXZlc1wiO1xuaW1wb3J0IEF1dG9JbXBvcnQgZnJvbSBcInVucGx1Z2luLWF1dG8taW1wb3J0L3ZpdGVcIjtcbmltcG9ydCBDb21wb25lbnRzIGZyb20gXCJ1bnBsdWdpbi12dWUtY29tcG9uZW50cy92aXRlXCI7XG5pbXBvcnQgeyBOYWl2ZVVpUmVzb2x2ZXIgfSBmcm9tIFwidW5wbHVnaW4tdnVlLWNvbXBvbmVudHMvcmVzb2x2ZXJzXCI7XG5pbXBvcnQgeyBBcmNvUmVzb2x2ZXIgfSBmcm9tIFwidW5wbHVnaW4tdnVlLWNvbXBvbmVudHMvcmVzb2x2ZXJzXCI7XG5cbi8vIGh0dHBzOi8vdml0ZWpzLmRldi9jb25maWcvXG5leHBvcnQgZGVmYXVsdCBkZWZpbmVDb25maWcoe1xuICBiYXNlOiBcIi9wdXp6bGVcIixcbiAgcmVzb2x2ZToge1xuICAgIGFsaWFzOiB7XG4gICAgICBcIkBcIjogcmVzb2x2ZShfX2Rpcm5hbWUsIFwiLi9zcmNcIiksXG4gICAgICBhcGk6IHJlc29sdmUoX19kaXJuYW1lLCBcIi4vc3JjL2FwaS9tZXRob2RzXCIpLFxuICAgIH0sXG4gIH0sXG5cbiAgcGx1Z2luczogW1xuICAgIHZ1ZSh7XG4gICAgICB0ZW1wbGF0ZToge1xuICAgICAgICBjb21waWxlck9wdGlvbnM6IHtcbiAgICAgICAgICBpc0N1c3RvbUVsZW1lbnQ6ICh0YWcpID0+IHRhZy5zdGFydHNXaXRoKFwiZGVmby1cIiksXG4gICAgICAgIH0sXG4gICAgICB9LFxuICAgIH0pLFxuICAgIHZ1ZUpzeCgpLFxuICAgIFVub0NTUyh7XG4gICAgICB0cmFuc2Zvcm1lcnM6IFt0cmFuc2Zvcm1lckRpcmVjdGl2ZSgpXSxcbiAgICAgIHByZXNldHM6IFtwcmVzZXRJY29ucyh7fSldLFxuICAgIH0pLFxuICAgIEF1dG9JbXBvcnQoe1xuICAgICAgaW1wb3J0czogW1xuICAgICAgICBcInZ1ZVwiLFxuICAgICAgICB7XG4gICAgICAgICAgXCJuYWl2ZS11aVwiOiBbXG4gICAgICAgICAgICBcInVzZURpYWxvZ1wiLFxuICAgICAgICAgICAgXCJ1c2VNZXNzYWdlXCIsXG4gICAgICAgICAgICBcInVzZU5vdGlmaWNhdGlvblwiLFxuICAgICAgICAgICAgXCJ1c2VMb2FkaW5nQmFyXCIsXG4gICAgICAgICAgXSxcbiAgICAgICAgfSxcbiAgICAgIF0sXG4gICAgICByZXNvbHZlcnM6IFtOYWl2ZVVpUmVzb2x2ZXIoKSwgQXJjb1Jlc29sdmVyKCldLFxuICAgIH0pLFxuICAgIENvbXBvbmVudHMoe1xuICAgICAgcmVzb2x2ZXJzOiBbXG4gICAgICAgIE5haXZlVWlSZXNvbHZlcigpLFxuICAgICAgICBBcmNvUmVzb2x2ZXIoe1xuICAgICAgICAgIHNpZGVFZmZlY3Q6IHRydWUsXG4gICAgICAgIH0pLFxuICAgICAgXSxcbiAgICB9KSxcbiAgXSxcbn0pO1xuIl0sCiAgIm1hcHBpbmdzIjogIjtBQUFrUixTQUFTLG9CQUFvQjtBQUMvUyxPQUFPLFNBQVM7QUFDaEIsU0FBUyxlQUFlO0FBQ3hCLE9BQU8sWUFBWTtBQUNuQixPQUFPLFlBQVk7QUFDbkIsT0FBTyxpQkFBaUI7QUFDeEIsT0FBTywwQkFBMEI7QUFDakMsT0FBTyxnQkFBZ0I7QUFDdkIsT0FBTyxnQkFBZ0I7QUFDdkIsU0FBUyx1QkFBdUI7QUFDaEMsU0FBUyxvQkFBb0I7QUFWN0IsSUFBTSxtQ0FBbUM7QUFhekMsSUFBTyxzQkFBUSxhQUFhO0FBQUEsRUFDMUIsTUFBTTtBQUFBLEVBQ04sU0FBUztBQUFBLElBQ1AsT0FBTztBQUFBLE1BQ0wsS0FBSyxRQUFRLGtDQUFXLE9BQU87QUFBQSxNQUMvQixLQUFLLFFBQVEsa0NBQVcsbUJBQW1CO0FBQUEsSUFDN0M7QUFBQSxFQUNGO0FBQUEsRUFFQSxTQUFTO0FBQUEsSUFDUCxJQUFJO0FBQUEsTUFDRixVQUFVO0FBQUEsUUFDUixpQkFBaUI7QUFBQSxVQUNmLGlCQUFpQixDQUFDLFFBQVEsSUFBSSxXQUFXLE9BQU87QUFBQSxRQUNsRDtBQUFBLE1BQ0Y7QUFBQSxJQUNGLENBQUM7QUFBQSxJQUNELE9BQU87QUFBQSxJQUNQLE9BQU87QUFBQSxNQUNMLGNBQWMsQ0FBQyxxQkFBcUIsQ0FBQztBQUFBLE1BQ3JDLFNBQVMsQ0FBQyxZQUFZLENBQUMsQ0FBQyxDQUFDO0FBQUEsSUFDM0IsQ0FBQztBQUFBLElBQ0QsV0FBVztBQUFBLE1BQ1QsU0FBUztBQUFBLFFBQ1A7QUFBQSxRQUNBO0FBQUEsVUFDRSxZQUFZO0FBQUEsWUFDVjtBQUFBLFlBQ0E7QUFBQSxZQUNBO0FBQUEsWUFDQTtBQUFBLFVBQ0Y7QUFBQSxRQUNGO0FBQUEsTUFDRjtBQUFBLE1BQ0EsV0FBVyxDQUFDLGdCQUFnQixHQUFHLGFBQWEsQ0FBQztBQUFBLElBQy9DLENBQUM7QUFBQSxJQUNELFdBQVc7QUFBQSxNQUNULFdBQVc7QUFBQSxRQUNULGdCQUFnQjtBQUFBLFFBQ2hCLGFBQWE7QUFBQSxVQUNYLFlBQVk7QUFBQSxRQUNkLENBQUM7QUFBQSxNQUNIO0FBQUEsSUFDRixDQUFDO0FBQUEsRUFDSDtBQUNGLENBQUM7IiwKICAibmFtZXMiOiBbXQp9Cg==
