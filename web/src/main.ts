import { createApp } from 'vue'
import App from './App.vue'
import { pinia } from "@/store";
import router from "@/router/router.ts";
import Antd from 'ant-design-vue';

import 'ant-design-vue/dist/reset.css';
import "@/styles/index.css";
import 'virtual:uno.css'

createApp(App).use(Antd).use(router).use(pinia).mount("#app");
