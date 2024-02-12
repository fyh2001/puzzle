import { defineStore } from "pinia";
import { authRequest } from "api/admin";
import type { AuthReq } from "@/types/admin";

export const useAdminStore = defineStore("admin", {
  state: () => ({
    authorization: {
      token: "",
      secretUrl: "",
      expiredAt: 0,
    },
  }),

  actions: {
    async auth(data: AuthReq): Promise<boolean> {
      const {
        data: { code, data: authResp },
      } = await authRequest.authorization(data);

      if (code === 200) {
        this.authorization.token = authResp.token;
        this.authorization.expiredAt = Date.now() + 24 * 60 * 60 * 1000; // 24小时
        return true;
      }

      return false;
    },
  },

  persist: [
    {
      paths: ["authorization"],
      storage: localStorage,
    },
  ],
});
