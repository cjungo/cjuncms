import { defineStore } from "pinia";
import { reactive, ref } from "vue";

export const useAuthStore = defineStore('auth', () => {
    const token = ref("");
    const permissions: any[] = reactive([]);

    return {
        token,
        permissions,
    };
}, {
    persist: true,
});