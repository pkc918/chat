import { defineConfig } from 'vite'
import Vue from '@vitejs/plugin-vue'
import UnoCSS from "unocss/vite";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [Vue(), UnoCSS()],
  server:{
    proxy: {
      "/api": {
        target: 'http://127.0.0.1:8080/api',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      }
    }
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src")
    }
  }
})
