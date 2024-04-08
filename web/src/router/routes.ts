import { RouteRecordRaw } from "vue-router";

export const routes: RouteRecordRaw[] = [{
    path: "/chat",
    name: "Chat",
    component: () => import("@/view/Chat.vue")
}]
