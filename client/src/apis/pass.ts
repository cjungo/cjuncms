import { apiGet } from "../utils/api";

export type CjPass = {
  id: number; // ID
  type: number; // 0.密码；1.密钥
  title: string; // 名称
  host: string; // 主机
  port: number; // 端口
  content: string; // 内容
};

export type QueryPassParam = {
  skip: number;
  take: number;
  plain: string;
};
export const queryPass = apiGet<QueryPassParam, CjPass[]>("/api/pass/query");
