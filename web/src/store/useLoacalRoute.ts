import { defineStore } from "pinia";
import { ref } from "vue";
import { RouteRecordName } from "vue-router";

export const useLocalRoute = defineStore("localRoute", () => {
    const routeName = ref<RouteRecordName>("");

    function setRouteName(pathName: RouteRecordName) {
        routeName.value = pathName;
    }

    return {routeName, setRouteName};
});

