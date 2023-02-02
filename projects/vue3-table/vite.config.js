import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import {resolve} from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: [//配置别名
      { find: '@', replacement: resolve(__dirname, 'src') }
    ],
    // 情景导出 package.json 配置中的exports字段
    conditions: [],
    // 导入时想要省略的扩展名列表
    // 不建议使用 .vue 影响IDE和类型支持
    extensions:['.mjs','.js','.ts','.jsx','.tsx','.json']
  },
})
