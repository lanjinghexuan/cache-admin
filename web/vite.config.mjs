import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 3000,
    proxy: {
      '/getUser': 'http://localhost:8080',
      '/cache': 'http://localhost:8080'
    }
  }
}) 