import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueRouter from 'vue-router/vite'

export default defineConfig(({ command }) => ({
  plugins: [
    vueRouter({
      routesFolder: 'src/pages',
      dts: 'src/typed-router.d.ts',
      watch: command === 'serve',
    }),
    vue(),
  ],
  resolve: {
    alias: {
      '~': fileURLToPath(new URL('./src', import.meta.url)),
    },
  },
  server: {
    host: '127.0.0.1',
    port: 3333,
    proxy: {
      '/_ssr/data': {
        target: 'http://127.0.0.1:4001',
        // Keep the public Vite origin for gossr's same-origin check.
        changeOrigin: false,
      },
    },
  },
  build: {
    target: 'es2020',
  },
}))
