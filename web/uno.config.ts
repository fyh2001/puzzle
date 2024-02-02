import { defineConfig, presetAttributify, presetUno } from 'unocss'

export default defineConfig({
  presets: [presetAttributify(), presetUno()],
  content: {
    pipeline: {
      include: [
        // 默认设置
        /\.(vue|svelte|[jt]sx|mdx?|astro|elm|php|phtml|html)($|\?)/,
        // 项目中的文件
        'src/**/*.{js,ts}',
      ],
    },
  },
  preflights: [
    {
      getCSS: () => `
      .el-button {
        background-color: var(--el-button-bg-color, var(--el-color-white))
      }
      .n-button {
        background-color: var(--n-color)
      }
      `,
    },
  ],
})
