import { apiDelete, apiPost } from "../utils/api";

export type CjDepartment = {
  id: number; //
  parent_id: number; //
  title: string; // 名称
  create_at: string; //
};

export type CjDepartmentEmployee = {
  department_id: number; //
  employee_id: number; //
  position_id: number; //
  create_at: string; //
};

export type CjDepartmentPosition = {
  id: number; //
  department_id: number; // 部门ID
  title: string; //
  create_at: string; //
};

export type QueryDepartmentParam = {};

export type DepartmentNode = CjDepartment & {
  children?: Array<DepartmentNode>;
};

export const queryDepartment = apiPost<
  QueryDepartmentParam,
  Array<DepartmentNode>
>("/api/department/query");
