import { defineStore } from "pinia";
import { reactive } from "vue";

export const useAppStore = defineStore('app', () => {
    const dialog = reactive({
        logout: {
            visible: false,
        },
    });
    const tabbar = reactive({
        items: [],
    } as {
        items: any[],
    });
    const navbar = reactive({
        opends: [],
    } as {
        opends: string[],
    });

    return {
        dialog,
        tabbar,
        navbar,
    };
}, {
    persist: true,
});