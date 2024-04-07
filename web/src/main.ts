import { createApp } from 'vue'
import App from './App.vue'
import router from "@/router/router.ts";
import "@/styles/index.css";
import 'virtual:uno.css'

createApp(App).use(router).mount("#app");
