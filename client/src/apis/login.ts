import { apiGet, apiPost } from "../utils/api";
import { CjEmployee } from "./employee";

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

export type Profile = CjEmployee & {};

export const login = apiPost<LoginParam, Sign>("/sign/in");
export const logout = apiPost<any, any>("/sign/out");
export const renewal = apiGet<any, string>("/sign/renewal");
export const getProfile = apiGet<any, Profile>("/sign/profile");
