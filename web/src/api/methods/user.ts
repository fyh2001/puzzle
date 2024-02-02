import { request } from "@/api/index";
import type { UserModel } from "@/types/user";

const baseURL = "/user";

export const userRequest = {
  register: (data: UserModel) =>
    request({
      url: `${baseURL}/register`,
      method: "POST",
      data,
    }),
};
