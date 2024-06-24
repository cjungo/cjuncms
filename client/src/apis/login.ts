import { apiGet, apiPost } from "../utils/api";

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

export type Sign = {
  token: string;
  permissions: [string];
  user: User;
};

export const login = apiPost<LoginParam, Sign>("/sign/in");
export const renewal = apiGet<any, string>("/sign/renewal");