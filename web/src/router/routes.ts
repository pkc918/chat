import { RouteRecordRaw } from "vue-router";

export const routes: RouteRecordRaw[] = [{
    path: "/",
    redirect: {
        name: "Page"
    }
}, {
    path: "/page",
    name: "Page",
    redirect: {
        name: "Chat"
    },
    component: () => import("@/components/Page/Page.vue"),
    children: [{
        path: "chat",
        name: "Chat",
        component: () => import("@/view/Chat.vue")
    }, {
        path: "friend",
        name: "Friend",
        component: () => import("@/view/Friends.vue")
    }, {
        path: "collect",
        name: "Collect",
        component: () => import("@/view/Collect.vue")
    }, {
        path: "comment",
        name: "Comment",
        component: () => import("@/view/Comment.vue")
    }, {
        path: "battery",
        name: "Battery",
        component: () => import("@/view/Battery.vue")
    }, {
        path: "weather",
        name: "Weather",
        component: () => import("@/view/Comment.vue")
    }, {
        path: "calendar",
        name: "Calendar",
        component: () => import("@/view/Calendar.vue")
    }, {
        path: "setting",
        name: "Setting",
        component: () => import("@/view/Setting.vue")
    }]
}, {
    path: "/logIn",
    name: "LogIn",
    component: () => import("@/view/LogIn.vue")
}];
