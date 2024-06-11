import { defineStore } from "pinia";
import { reactive, ref } from "vue";

export type User = {
  id: number;
  fullname?: string;
  nickname?: string;
  avatar_path?: string;
};

export const useAuthStore = defineStore(
  "auth",
  () => {
    const token = ref("");
    const permissions: any[] = reactive([]);
    const user = reactive<User>({ id: 0 });

    return {
      token,
      permissions,
      user,
    };
  },
  {
    persist: true,
  }
);
