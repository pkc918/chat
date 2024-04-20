import { createRouter, createWebHistory, RouteLocationNormalized } from "vue-router";
import { routes } from "@/router/routes.ts";
import { useLocalRoute } from "@/store/useLoacalRoute.ts";

const router = createRouter({
    history: createWebHistory(),
    routes: routes
});

router.beforeEach((to: RouteLocationNormalized, from: RouteLocationNormalized) => {
    console.log("to", to);
    console.log("from", from);
    if (to.name) {
        useLocalRoute().setRouteName(to.name!);
    }
    return true;
});

export default router
