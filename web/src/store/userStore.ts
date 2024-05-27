import { defineStore } from "pinia";
import { ref } from "vue";

export const useUserStore = defineStore("userStore", () => {
    const token = ref<string>("");

    const setToken = (newToken: string) => {
        token.value = newToken;
    };

    const clearToken = () => {
        token.value = "";
    }

    return {token, setToken, clearToken};
}, {
    persist: true
})

