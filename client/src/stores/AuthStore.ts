import { defineStore } from "pinia";
import { reactive, ref } from "vue";
import { type User } from "../apis/login";

export const useAuthStore = defineStore(
  "auth",
  () => {
    const token = ref("");
    const permissions: string[] = reactive([]);
    const user = reactive<User>({ id: 0, username: "" });

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
