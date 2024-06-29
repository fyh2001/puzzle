import { defineStore } from "pinia";
import { userRequest } from "api/user";
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

    async fetchUser() {
      const {
        data: { code, data: userInfo },
      } = await userRequest.getUserInfo();

      if (code === 200) {
        this.setUser(userInfo);
      }
    },
  },
});
