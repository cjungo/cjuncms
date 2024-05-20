import { type AxiosResponse } from "axios";
import { ApiResult, api } from "../utils/api";

export type LoginParam = {
  username: string;
  password: string;
  captchaId: string;
  captchaAnswer: string;
};

export type LoginResult = ApiResult<{
  token: string;
}>;

export const login = async (
  param: LoginParam
): Promise<AxiosResponse<LoginResult>> => {
  return await api.post("/login", param);
};
