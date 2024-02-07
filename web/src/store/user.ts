import { defineStore } from "pinia";
import type { UserResp } from "@/types/user";

export const useUserStore = defineStore("user", {
  persist: true,

  state: () => ({
    user: <UserResp>{
      id: "",
      username: "",
      nickname: "",
      email: "",
      phone: "",
      avatar: "",
      status: 0,
      createdAt: "",
      updatedAt: "",
    },
    token: <string | null>null,
  }),

  getters: {
    getUser: (state) => state.user,
    getToken: (state) => state.token,
  },

  actions: {
    setUser(user: UserResp) {
      this.user = user;
    },

    setToken(token: string) {
      this.token = token;
    },

    logout() {
      this.user = {
        id: "",
        username: "",
        nickname: "",
        email: "",
        phone: "",
        avatar: "",
        status: 0,
        createdAt: "",
        updatedAt: "",
      };
      this.token = null;
    },
  },
});
