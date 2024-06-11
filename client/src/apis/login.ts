import { type AxiosResponse } from "axios";
import { ApiResult, api } from "../utils/api";

export type User = {
  username: string;
  nickname?: string;
};

export type LoginParam = {
  username: string;
  password: string;
  captchaId: string;
  captchaAnswer: string;
};

export type LoginResult = ApiResult<{
  token: string;
  permissions?: [string];
  user?: User;
}>;

export const login = async (
  param: LoginParam
): Promise<AxiosResponse<LoginResult>> => {
  return await api.post("/sign/in", param);
};
