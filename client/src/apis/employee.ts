import { apiDelete, apiPost } from "../utils/api";

export type CjEmployee = {
  id: number; // ID
  jobnumber: string; // 工号
  username: string; // 用户名
  nickname?: string; // 昵称
  fullname?: string; // 全称
  birthday?: string; // 生日
  avatar_path?: string; // 头像路径
  is_removed: number; // 是否删除
};

export type QueryEmployeeParam = {
  skip: number;
  take: number;
  plain: string;
};

export const queryEmployee = apiPost<QueryEmployeeParam, CjEmployee[]>(
  "/api/employee/query"
);

export type AddEmployeeParam = Omit<CjEmployee, "id">;
export const addEmployee = apiPost<AddEmployeeParam, CjEmployee>(
  "/api/employee/add"
);

export type EditEmployeeParam = CjEmployee;
export const editEmployee = apiPost<EditEmployeeParam, any>("api/employee/edit");

export type DropEmployeeParam = {
  id?: number;
  ids?: number[];
};
export const dropEmployee = apiDelete<DropEmployeeParam, any>(
  "/api/employee/drop"
);
