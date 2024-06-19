import { ApiResult, api } from "../utils/api";

export type User = {
  id: number;
  username: string;
  fullname?: string;
  nickname?: string;
  avatar_path?: string;
};

export type LoginParam = {
  username: string;
  password: string;
  captchaId: string;
  captchaAnswer: string;
};

export type LoginResult = ApiResult<{
  token: string;
  permissions: [string];
  user: User;
}>;

export const login = async (param: LoginParam): Promise<LoginResult> => {
  return await api.post("/sign/in", param);
};
