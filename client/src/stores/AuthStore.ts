import { defineStore } from "pinia";
import { ref } from "vue";
import { type User } from "../apis/login";

export const useAuthStore = defineStore(
  "auth",
  () => {
    const token = ref("");
    const permissions = ref<string[]>([]);
    const user = ref<User>({ id: 0, username: "" });

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
