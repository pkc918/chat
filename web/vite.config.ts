import { defineConfig } from 'vite'
import Vue from '@vitejs/plugin-vue'
import UnoCSS from "unocss/vite";
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [Vue(), UnoCSS()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src")
    }
  }
})
