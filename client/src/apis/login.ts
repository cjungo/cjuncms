import { api } from "../utils/api";

export type LoginParam = {
  username: string;
  password: string;
  captchaId: string;
  captchaAnswer: string;
};

export type LoginResult = {
  code: number;
  data: string;
};

export const login = async (param: LoginParam): Promise<LoginResult> => {
  const response = await api.post("/login", param);
  return response.data
};
