import { apiDelete, apiPost } from "../utils/api";

export type CjScript = {
  id: number; // ID
  pass_id: number; // 密钥ID
  project_id: number; // 项目ID
  title: string; // 标题
  content: string; // 内容
};

export type QueryScriptParam = {};
export const queryScript = apiPost<QueryScriptParam, CjScript[]>(
  "/api/script/query"
);

export type AddScriptParam = CjScript;
export const addScript = apiPost<AddScriptParam, CjScript>("/api/script/add");

export type EditScriptParam = CjScript;
export const editScript = apiPost<EditScriptParam, CjScript>(
  "/api/script/edit"
);

export type DropScriptParam = {
  id?: number;
  ids?: number[];
};
export const dropScript = apiDelete<DropScriptParam, any>("/api/script/drop");
